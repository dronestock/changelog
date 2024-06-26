FROM quay.io/git-chglog/git-chglog:0.15.4 AS chglog




FROM ccr.ccs.tencentyun.com/storezhang/alpine:3.20.0


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成Changelog插件，主要是用来生成改变日志，并结合发布插件发布产品"


# 复制文件
COPY docker /
COPY changelog /bin
COPY --from=chglog /usr/local/bin/git-chglog /usr/local/bin/git-chglog


RUN set -ex \
    \
    \
    \
    && apk update \
    \
    # 改变日志依赖于Git查询提交记录
    && apk --no-cache add git \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/changelog \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 默认参数
ENV PLUGIN_CONFIG "expr: file(\"/etc/changelog/config.yml.gohtml\")"
ENV PLUGIN_TEMPLATE "expr: file(\"/etc/changelog/CHANGELOG.md.gohtml\")"
ENV CONFIG_PATH /tmp/config.yml
ENV TEMPLATE_PATH /tmp/CHANGELOG.tpl.md


# 执行命令
ENTRYPOINT /bin/changelog
