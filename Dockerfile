FROM storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成Changelog插件，主要是用来生成改变日志，并结合发布插件发布产品"


# 复制文件
COPY docker /
COPY changelog /bin


RUN set -ex \
    \
    \
    \
    && apk update \
    \
    # 改变日志依赖于Git查询提交记录
    && apk --no-cache add git \
    \
    && apk --no-cache add npm \
    \
    # 配置镜像加速安装过程
    && npm config set registry http://registry.npmmirror.com \
    && npm install -g standard-version \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/changelog \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/changelog
