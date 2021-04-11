package proxy_test

import (
	"fmt"
	"structural/proxy"
	"testing"
)

// Test
func Test(t *testing.T) {
	// create object and proxy
	mo := proxy.NewMyObject("Hello, world")
	p := proxy.NewProxy(mo)

	// use proxy functions
	obVal := mo.Val()
	pVal := p.Val()

	if obVal != "Hello, world" {
		t.Errorf("obVal is initially %s and should be Hello, world\n", obVal)
	}

	if pVal != "Hello, world as found by the proxy" {
		t.Errorf("pVal is initially %s and should be Hello, world as found by the proxy\n", obVal)
	}

	mo.Apply("Goodbye")
	obVal = mo.Val()
	pVal = p.Val()

	if obVal != "Goodbye" {
		t.Errorf("obVal is %s and should be Goodbye\n", obVal)
	}

	if pVal != "Goodbye as found by the proxy" {
		t.Errorf("pVal is %s and should be Goodbye as found by the proxy\n", obVal)
	}

	p.Apply("Goodbye Again")
	obVal = mo.Val()
	pVal = p.Val()

	if obVal != "Proxy set val Goodbye Again" {
		t.Errorf("obVal is %s and should be Goodbye Again\n", obVal)
	}

	if pVal != "Proxy set val Goodbye Again as found by the proxy" {
		t.Errorf("pVal is %s and should be Proxy set val Goodbye Again as found by the proxy\n", obVal)
	}

}

// Example Test
func Example() {

	// create object and proxy
	mo := proxy.NewMyObject("Hello, world")
	p := proxy.NewProxy(mo)

	// use proxy functions
	fmt.Printf("Getting the value from MyObject: %s\n", mo.Val())
	fmt.Printf("Getting the value from Proxy: %s\n", p.Val())

	mo.Apply("Goodbye")
	fmt.Printf("Getting the value from MyObject: %s\n", mo.Val())
	fmt.Printf("Getting the value from Proxy: %s\n", p.Val())

	p.Apply("Goodbye Again")
	fmt.Printf("Getting the value from MyObject: %s\n", mo.Val())
	fmt.Printf("Getting the value from Proxy: %s\n", p.Val())

	//Output:Getting the value from MyObject: Hello, world
	//Getting the value from Proxy: Hello, world as found by the proxy
	//Getting the value from MyObject: Goodbye
	//Getting the value from Proxy: Goodbye as found by the proxy
	//Getting the value from MyObject: Proxy set val Goodbye Again
	//Getting the value from Proxy: Proxy set val Goodbye Again as found by the proxy

}
