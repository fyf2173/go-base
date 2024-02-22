一款golang的项目开发目录结构，简洁明了，非常适模块化业务的开发。

```sh
./go-base
├─cmd
├─internal
│  ├─middleware
│  ├─modules
│  │  ├─module_entry.go
│  │  ├─example
│  │  │  ├─service.go
│  │  │  ├─api
│  │  │  │  ├─handle.go
│  │  │  │  └─route.go
│  │  │  └─http
│  │  └─wsocket
│  │      └─http
│  ├─orm
│  └─pkg
│      ├─common
│      └─job
└─resources
```
