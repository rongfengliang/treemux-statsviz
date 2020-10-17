#  statsviz for treemux


## how to using


```code
go get github.com/rongfengliang/treemux-statsviz
router := treemux.New()
pprof.RouterStatsvizRegister(router)

view result:

/debug/statsviz

```