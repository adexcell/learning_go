package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Crawler config & state
type Crawler struct {
	OutputDirectory string
	MaxDepth        int

	visited map[string]bool
	mu      sync.Mutex

	// Concurrency primitives
	wg     sync.WaitGroup
	tokens chan struct{}
}

// NewCrawler creates a new instance with initialized maps and channels
func NewCrawler(maxConcurrency int, outputDir string) *Crawler {
	return &Crawler{
		OutputDirectory: outputDir,
		visited:         make(map[string]bool),
		tokens:          make(chan struct{}, maxConcurrency),
	}
}

// Run starts the crawling process
func (c *Crawler) Run(startURL string, depth int) {
	c.MaxDepth = depth
	
	c.wg.Add(1)
	go c.visit(startURL, 0)

	c.wg.Wait()
	fmt.Println("Done! All tasks finished.")
}

func (c *Crawler) visit(currentURL string, depth int) {
	defer c.wg.Done()

	if depth > c.MaxDepth {
		return
	}

	c.mu.Lock()
	if c.visited[currentURL] {
		c.mu.Unlock()
		return
	}
	c.visited[currentURL] = true
	c.mu.Unlock()

	c.tokens <- struct{}{}
	
	links, err := c.processPage(currentURL)
	
	<-c.tokens

	if err != nil {
		log.Printf("Error processing %s: %v", currentURL, err)
		return
	}

	for _, link := range links {
		c.wg.Add(1)
		go c.visit(link, depth+1)
	}
}

// processPage downloads, saves, and parses a single page (Worker)
func (c *Crawler) processPage(targetURL string) ([]string, error) {
	fmt.Printf("Processing: %s\n", targetURL)

	body, contentType, err := c.fetch(targetURL)
	if err != nil {
		return nil, err
	}

	var links []string
	var finalBody []byte

	if strings.Contains(contentType, "text/html") {
		finalBody, links, err = c.rewriteLinks(body, targetURL)
		if err != nil {
			return nil, err
		}
	} else {
		finalBody = body
		links = nil
	}

	if err := c.saveFile(targetURL, finalBody); err != nil {
		return nil, err
	}

	baseURL, _ := url.Parse(targetURL)
	var validLinks []string
	for _, link := range links {
		parsed, err := url.Parse(link)
		if err != nil {
			continue
		}
		if parsed.Host == baseURL.Host {
			validLinks = append(validLinks, parsed.String())
		}
	}

	return validLinks, nil
}

// fetch performs the HTTP GET request
func (c *Crawler) fetch(targetURL string) ([]byte, string, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, "", fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return body, resp.Header.Get("Content-Type"), nil
}

// rewriteLinks parses HTML, rewrites href/src to relative paths, and returns found links
func (c *Crawler) rewriteLinks(body []byte, currentURL string) ([]byte, []string, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, nil, fmt.Errorf("parse error: %w", err)
	}

	baseURL, err := url.Parse(currentURL)
	if err != nil {
		return nil, nil, err
	}

	currentFilePath := c.urlToPath(currentURL)
	currentDir := filepath.Dir(currentFilePath)

	var links []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attrKey string
			switch n.Data {
			case "a", "link":
				attrKey = "href"
			case "img", "script":
				attrKey = "src"
			}

			if attrKey != "" {
				for i, attr := range n.Attr {
					if attr.Key == attrKey {
						parsedVal, err := url.Parse(attr.Val)
						if err == nil {
							absLink := baseURL.ResolveReference(parsedVal)
							links = append(links, absLink.String())

							targetPath := c.urlToPath(absLink.String())
							relPath, err := filepath.Rel(currentDir, targetPath)
							if err == nil {
								n.Attr[i].Val = filepath.ToSlash(relPath)
							}
						}
						break
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(doc)

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), links, nil
}

// saveFile saves content to the OutputDirectory
func (c *Crawler) saveFile(u string, content []byte) error {
	relPath := c.urlToPath(u)
	finalPath := filepath.Join(c.OutputDirectory, relPath)

	if err := os.MkdirAll(filepath.Dir(finalPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(finalPath, content, 0644)
}

// urlToPath helper to convert URL to local relative path
func (c *Crawler) urlToPath(u string) string {
	parsed, err := url.Parse(u)
	if err != nil {
		return "invalid_url"
	}

	path := parsed.Path
	if path == "" || strings.HasSuffix(path, "/") {
		path += "index.html"
	} else if !strings.Contains(filepath.Base(path), ".") {
		path += ".html"
	}
	
	path = strings.TrimPrefix(path, "/")

	return filepath.Join(parsed.Host, path)
}

func main() {
	depth := flag.Int("l", 0, "Recursion depth")
	concurrency := flag.Int("c", 10, "Max concurrency")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run main.go [-l depth] [-c concurrency] <url>")
		os.Exit(1)
	}

	crawler := NewCrawler(*concurrency, "downloads")

	for _, url := range args {
		crawler.Run(url, *depth)
	}
}