# daud-sdk

A simple Go SDK that exposes public profile information for **Daud Ahmad**.

## Installation

```bash
go get github.com/daud/daud-sdk@v0.1.0
```

## Usage

```bash
package main

import (
	"fmt"

	"github.com/daud/daud-sdk"
)

func main() {
	profile := daudsdk.GetProfile()

	fmt.Println(profile.Name)
	fmt.Println(profile.GitHub)
	fmt.Println(profile.LinkedIn)
	fmt.Println(profile.Website)
}
```
---

Checkout my Profiles:

[Website](https://daudahmad.com) | [Linkedin](https://www.linkedin.com/in/daudahmad0303/) | [Github](https://github.com/DaudAhmad0303)

If above doesn't render
<p>
  <a href="https://daudahmad.com">https://daudahmad.com</a> |
  <a href="https://www.linkedin.com/in/daudahmad0303/">https://www.linkedin.com/in/daudahmad0303/</a> |
  <a href="https://github.com/DaudAhmad0303">https://github.com/DaudAhmad0303</a>
</p>

---

