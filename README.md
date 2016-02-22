# Doploy
`doploy` is simple cli for updating Marathon apps in Go.


## Getting Started

``` go
go install github.com/yuzu-co/doploy

```

## Usage

Set your Marathon API Url in env
```
export MARATHON_URL=https://user@pass:your-marathon
```


``` go
#update an app
doploy update [service_name] --scale=1 --mem=200 --cpu=2

#update an app, exit only when done, usefull for continuous deployment
doploy update [service_name] --scale=1 --mem=200 --cpu=2 --sync=1

#upgrade an image/tag with a public image
doploy update [service_name] --image=mongo:3.1.8

#upgrade an image/tag from a private registry
doploy update [service_name] --image=registry.company.com:5000/namespace/image:tag

```

## Todo

- add tests
