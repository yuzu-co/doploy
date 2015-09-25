# Doploy
`doploy` is simple cli for updating Marathon apps in Go.


## Getting Started

``` go
go install github.com/yuzu-co/doploy

```

## Usage

Set your Marathon API Url in env
```
export MARATHON_URL=https://user@pass:your-marathon/v2/
```


``` go
#update an app
doploy update [service_name] --scale=1 --mem=200 --cpu=2

#upgrade an image/tag with a public image
doploy update [service_name] --image=mongo:3.1.8

#upgrade an image/tag with a private registry
doploy update [service_name] --image=registry.company.com:5000/namespace/image:tag

```

## Todo

- add tests