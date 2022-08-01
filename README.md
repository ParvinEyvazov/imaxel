# imaxel

imaxel is a image operation project for Web that written with using Golang & Webassembly.

## Services

1 - Find pixel differences between 2 images and report back different ones (pixel).

## How to run the project?

- First clone the repo

```
git clone https://github.com/ParvinEyvazov/imaxel.git
```

- Setup (will enable wasm_exec.js file)

```
make setup
```

- Build the project (Compile .wasm file with the updated cmd/wasm/main.go)

```
make build
```

- Run the project (Running on port Listening on http://localhost:3000)

```
make run
```

OR

- You can build and run in the same time

```
make br
```

Time to get into the code.

- After editing the code, for commiting:

```
make unsetup
```

- Send Pull Request

## Roadmap

- [ ] Pixel difference between 2 images

- [ ] Find color code on specific pixel or part of image
