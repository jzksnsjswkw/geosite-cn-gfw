# GeoSite CN + GFW

自动从 [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat) 提取 `CN` 和 `GFW` 两个分类，生成精简版 `GeoSite.dat`。

## 体积对比

| 版本 | 大小 | 分类数 |
|---|---|---|
| 完整版 | 10.9 MB | 1563 |
| **本仓库** | **1.9 MB** | **2 (CN + GFW)** |

## 下载

```
https://github.com/jzksnsjswkw/geosite-cn-gfw/releases/latest/download/GeoSite-cn-gfw.dat
```

## 用于 OpenClash

```bash
wget -O /etc/openclash/GeoSite.dat \
  https://github.com/jzksnsjswkw/geosite-cn-gfw/releases/latest/download/GeoSite-cn-gfw.dat
/etc/init.d/openclash restart
```

用于 Clash/Mihomo：
```yaml
geodata-mode: true
geox-url:
  geosite: "https://github.com/jzksnsjswkw/geosite-cn-gfw/releases/latest/download/GeoSite-cn-gfw.dat"
```

## 自动更新

- 每天 UTC 03:00（北京时间 11:00）自动构建
- 也可手动触发：Actions → Build GeoSite CN+GFW → Run workflow

## 内容

| 分类 | 域名数 | 用途 |
|---|---|---|
| `CN` | ~110,000 | 中国大陆域名（`GEOSITE,cn,DIRECT`） |
| `GFW` | ~4,400 | 被墙域名（`GEOSITE,gfw,PROXY`） |

## 注意事项

**只包含 CN 和 GFW 两个分类。** 如果你的配置用到其他分类（如 `google`、`youtube`、`netflix`、`category-ads-all`），需要改用完整版。

## 工具

`cmd/geosite-slim` 是一个通用的 GeoSite 裁剪工具：

```bash
go build -o geosite-slim ./cmd/geosite-slim
./geosite-slim geosite.dat out.dat CN GFW
./geosite-slim geosite.dat out.dat CN GFW GOOGLE YOUTUBE
```

## License

MIT
