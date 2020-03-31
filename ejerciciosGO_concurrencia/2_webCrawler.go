package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
   // Fetch devuelve el cuerpo de una URL y 
   // un slice de las URLs encontradas en la página
	Fetch(url string) (body string, urls []string, err error)
}


var crawledURLs = make(map[string]bool) /* registrador de urls utilizadas. */
var mux sync.Mutex						/* mutex para el control de hilos concurrentes. */

func CrawlURL(url string, depth int, fetcher Fetcher, quit chan bool) {
    // TODO: Obtener URLs en paralelo.
	// TODO: Evitar repeticiones.

    // En la primer instancia a la función, se bloquearía para siempre si nadie recibe desde el otro extremo de este canal
    defer func() { quit <- true }()

    if depth <= 0 {
        return
    }

    mux.Lock() 				/* exclusion mutua para ejecutar, paso solo un hilo a la vez. */
    _, isCrawled := crawledURLs[url]
    if isCrawled {			/* Si url fue utilizada termina ejecucion. */
        mux.Unlock()		/* Libero exclusion de ejecucion. Se permite ejecucion de otro hilo. */
        return
    }
    crawledURLs[url] = true /* indico la url como utilizada. */
    mux.Unlock()			/* Libero exclusion de ejecucion. Se permite ejecucion de otro hilo. */

    body, urls, err := fetcher.Fetch(url) /* traer url.*/
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    quitThis := make(chan bool)
    for _, u := range urls {
        go CrawlURL(u, depth-1, fetcher, quitThis)
    }
    for range urls {
        <-quitThis
    }
    return
}

// Crawl usa el 'fetcher' para obtener recursivamente las páginas
// empezando en cierta 'url', con cierta profundidad máxima 'depth'.
func Crawl(url string, depth int, fetcher Fetcher) {
    lastQuit := make(chan bool)
    go CrawlURL(url, depth, fetcher, lastQuit)
    
    // Se recibe finlizacion de gorutina desde este canal para desbloquear la función llamada
    <-lastQuit
    return
}


func main() {

	Crawl("http://golang.org/", 4, fetcher)
}


// fakeFetcher es un Fetcher the devuelve resultados ficticios.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls     []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher contiene los datos ficticios.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}