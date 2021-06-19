---
title: "Ubuntu上のdocker build で error: unpacking of archive failed on"
date: 2016-01-13T23:16:05+09:00
emoji: "📣"
type: "tech"
topics: [Docker, Ubuntu14.04, aufs]
published: true
---

## はじめに

ubuntu ホスト上で Docker の CentOS イメージ使って docker build したらエラー。
今後もハマることがある気がしたので備忘録。

基本的には下記の記事と同じような内容。
[CentOS7 + systemd + Docker でインフラ CI をやっていく - Qiita](http://qiita.com/udzura/items/fa93f262bbe036a1413e)

最後の「おわりに」に書きましたが、aufs を新しくして`CONFIG_AUFS_XATTR`を有効にするのが根源的な解法っぽい。（未確認）
今回は個人的に Docker を触って見てるだけなので、Docker のストレージドライバを aufs 以外にするだけの手軽に解決できる方法を選択しています。

ちなみに、boot2docker でも同じ問題が起こっていたそうですが、そちらは修正されている模様。
[Bump kernel and AUFS versions, update kernel config by pstengel · Pull Request #818 · boot2docker/boot2docker](https://github.com/boot2docker/boot2docker/pull/818)

**追記**：
Ubuntu のバージョンを 15.05 や 15.10 に上げればこの問題は解決されるようなので、Ubuntu のバージョンを上げるのが楽かもしれません。

## 環境

ホスト OS の Ubuntu 上で CentOS7 の Docker イメージを動かしています。

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=14.04
DISTRIB_CODENAME=trusty
DISTRIB_DESCRIPTION="Ubuntu 14.04.3 LTS"

[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% arch
x86_64

[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% uname -a
Linux kaakaa-System-Product-Name 3.13.0-74-generic #118-Ubuntu SMP Thu Dec 17 22:52:10 UTC 2015 x86_64 x86_64 x86_64 GNU/Linux

[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% docker version
Client:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:12:04 UTC 2015
 OS/Arch:      linux/amd64

Server:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:12:04 UTC 2015
 OS/Arch:      linux/amd64

[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% docker info
Containers: 29
Images: 307
Server Version: 1.9.1
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 365
 Dirperm1 Supported: false
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 3.13.0-74-generic
Operating System: Ubuntu 14.04.3 LTS
CPUs: 8
Total Memory: 23.24 GiB
Name: kaakaa-System-Product-Name
WARNING: No swap limit support
```

## 事象

下記のような Dockerfile から`docker build`したところ、エラーが発生。
Dockerfile の内容はこちらを参考にした：[Dockerfile を使いこなす(1) | Think IT（シンクイット）](https://thinkit.co.jp/story/2015/10/06/6453)

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% cat Dockerfile
FROM		centos:centos7.1.1503
MAINTAINER	kaakaa
ENV		container docker
RUN		yum swap -y fakesystemd systemd && yum clean all
RUN		yum install -y httpd && yum clean all
RUN		echo "Hello Apache." > /var/www/html/index.html
RUN		systemctl enable httpd
EXPOSE		80
```

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% docker build -f ./Dockerfile -t centos:c71apache01 --no-cache=true .
Sending build context to Docker daemon 2.048 kB
Step 1 : FROM centos:c71docker0001
 ---> 7e436e7a3b77
Step 2 : MAINTAINER kaakaa
 ---> Running in 8fdbd6688031
 ---> a0275d7d2d5b
Removing intermediate container 8fdbd6688031
Step 3 : ENV container docker
 ---> Running in 31b9611c1e1b
 ---> 943e6c11843b
Removing intermediate container 31b9611c1e1b
Step 4 : RUN yum swap -y fakesystemd systemd && yum clean all
 ---> Running in 02272406c14e
Loaded plugins: fastestmirror

...

<途中省略>

...

Dependencies Resolved

================================================================================
 Package             Arch          Version                    Repository   Size
================================================================================
Installing:
 httpd               x86_64        2.4.6-40.el7.centos        base        2.7 M
Installing for dependencies:
 apr                 x86_64        1.4.8-3.el7                base        103 k
 apr-util            x86_64        1.5.2-6.el7                base         92 k
 centos-logos        noarch        70.0.6-3.el7.centos        base         21 M
 httpd-tools         x86_64        2.4.6-40.el7.centos        base         82 k
 mailcap             noarch        2.1.41-2.el7               base         31 k

Transaction Summary
================================================================================
Install  1 Package (+5 Dependent packages)

Total download size: 24 M
Installed size: 31 M
Downloading packages:
--------------------------------------------------------------------------------
Total                                              356 kB/s |  24 MB  01:09
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : apr-1.4.8-3.el7.x86_64                                       1/6
  Installing : apr-util-1.5.2-6.el7.x86_64                                  2/6
  Installing : httpd-tools-2.4.6-40.el7.centos.x86_64                       3/6
  Installing : centos-logos-70.0.6-3.el7.centos.noarch                      4/6
  Installing : mailcap-2.1.41-2.el7.noarch                                  5/6
  Installing : httpd-2.4.6-40.el7.centos.x86_64                             6/6Error unpacking rpm package httpd-2.4.6-40.el7.centos.x86_64

error: unpacking of archive failed on file /usr/sbin/suexec: cpio: cap_set_file
error: httpd-2.4.6-40.el7.centos.x86_64: install failed
  Verifying  : httpd-tools-2.4.6-40.el7.centos.x86_64                       1/6
  Verifying  : apr-1.4.8-3.el7.x86_64                                       2/6
  Verifying  : mailcap-2.1.41-2.el7.noarch                                  3/6
  Verifying  : apr-util-1.5.2-6.el7.x86_64                                  4/6
  Verifying  : centos-logos-70.0.6-3.el7.centos.noarch                      5/6
  Verifying  : httpd-2.4.6-40.el7.centos.x86_64                             6/6

Dependency Installed:
  apr.x86_64 0:1.4.8-3.el7
  apr-util.x86_64 0:1.5.2-6.el7
  centos-logos.noarch 0:70.0.6-3.el7.centos
  httpd-tools.x86_64 0:2.4.6-40.el7.centos
  mailcap.noarch 0:2.1.41-2.el7

Failed:
  httpd.x86_64 0:2.4.6-40.el7.centos

Complete!
The command '/bin/sh -c yum install -y httpd && yum clean all' returned a non-zero code: 1
```

httpd のインストールで失敗しているよう。

```
error: unpacking of archive failed on file /usr/sbin/suexec: cpio: cap_set_file
error: httpd-2.4.6-40.el7.centos.x86_64: install failed
```

## 調査

[cap_set_file not permitted on aufs storage driver only · Issue #6980 · docker/docker](https://github.com/docker/docker/issues/6980)

Docker 本家の issue でも 2014 年から既知の問題。~~（未だに治らないのはなぜ？）~~ => aufs の`CONFIG_AUFS_XATTR`追加により解消されている模様

今回はホスト Ubuntu の場合で発生しましたが、boot2docker を使用している時も（issue の立てられた 2014 年頃は）発生するようです。

[CentOS7 + systemd + Docker でインフラ CI をやっていく - Qiita](http://qiita.com/udzura/items/fa93f262bbe036a1413e)
同じエラーについて触れている Qiita 記事。

> aufs の場合だけ systemd のアップグレードができないらしい。。 doot2docker なんかはデフォルトが aufs である。一時的に devicemapper を使うとかすれば良いかもしれない。

## 解決：Docker のストレージドライバを変更する

[docker のストレージドライバーを devicemapper に変更する＠Ubuntu 14.04 | Mazn.net](http://www.mazn.net/blog/2015/12/13/1545.html)
上記を参考に`/etc/default/docker`に `DOCKER_OPTS="--storage-driver=devicemapper"`を追加して、docker を restart。
※既存の Docker イメージ/コンテナは新しいストレージドライバで使用できなくなるため注意

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% cat /etc/default/docker
# Use DOCKER_OPTS to modify the daemon startup options.
DOCKER_OPTS="--storage-driver=devicemapper"

[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% sudo service docker restart
docker stop/waiting
docker start/running, process 6791
```

ストレージドライバーが更新されました。

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% docker info
Containers: 0
Images: 0
Server Version: 1.9.1
Storage Driver: devicemapper
 Pool Name: docker-8:17-5113301-pool
 Pool Blocksize: 65.54 kB
 Base Device Size: 107.4 GB
 Backing Filesystem: ext4
 Data file: /dev/loop0
 Metadata file: /dev/loop1
 Data Space Used: 1.821 GB
 Data Space Total: 107.4 GB
 Data Space Available: 48.64 GB
 Metadata Space Used: 1.479 MB
 Metadata Space Total: 2.147 GB
 Metadata Space Available: 2.146 GB
 Udev Sync Supported: true
 Deferred Removal Enabled: false
 Deferred Deletion Enabled: false
 Deferred Deleted Device Count: 0
 Data loop file: /var/lib/docker/devicemapper/devicemapper/data
 Metadata loop file: /var/lib/docker/devicemapper/devicemapper/metadata
 Library Version: 1.02.77 (2012-10-15)
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 3.13.0-74-generic
Operating System: Ubuntu 14.04.3 LTS
CPUs: 8
Total Memory: 23.24 GiB
Name: kaakaa-System-Product-Name
WARNING: No swap limit support
```

補足：boot2docker でのストレージドライバ変更方法
[The PkgFarmers — Fixing yum install on boot2docker](http://pkgfarm.tumblr.com/post/114104687791/fixing-yum-install-on-boot2docker)

httpd が無事インストールされ、`docker build`が成功するようになりました。

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% docker build -f ./Dockerfile -t centos:c71apache01 --no-cache=true .
Sending build context to Docker daemon 2.048 kB
Step 1 : FROM centos:centos7.1.1503
 ---> 173339447b7e
Step 2 : MAINTAINER kaakaa
 ---> Running in 9cedbb5714a1
 ---> e3bcc4967ffb
Removing intermediate container 9cedbb5714a1
Step 3 : ENV container docker
 ---> Running in 7af6b54d9459
 ---> 8b51ce6f53b7
Removing intermediate container 7af6b54d9459

...

<省略>

...

Dependencies Resolved

================================================================================
 Package             Arch          Version                    Repository   Size
================================================================================
Installing:
 httpd               x86_64        2.4.6-40.el7.centos        base        2.7 M
Installing for dependencies:
 apr                 x86_64        1.4.8-3.el7                base        103 k
 apr-util            x86_64        1.5.2-6.el7                base         92 k
 centos-logos        noarch        70.0.6-3.el7.centos        base         21 M
 httpd-tools         x86_64        2.4.6-40.el7.centos        base         82 k
 mailcap             noarch        2.1.41-2.el7               base         31 k

Transaction Summary
================================================================================
Install  1 Package (+5 Dependent packages)

Total download size: 24 M
Installed size: 31 M
Downloading packages:
--------------------------------------------------------------------------------
Total                                              803 kB/s |  24 MB  00:31
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : apr-1.4.8-3.el7.x86_64                                       1/6
  Installing : apr-util-1.5.2-6.el7.x86_64                                  2/6
  Installing : httpd-tools-2.4.6-40.el7.centos.x86_64                       3/6
  Installing : centos-logos-70.0.6-3.el7.centos.noarch                      4/6
  Installing : mailcap-2.1.41-2.el7.noarch                                  5/6
  Installing : httpd-2.4.6-40.el7.centos.x86_64                             6/6
  Verifying  : httpd-2.4.6-40.el7.centos.x86_64                             1/6
  Verifying  : httpd-tools-2.4.6-40.el7.centos.x86_64                       2/6
  Verifying  : apr-1.4.8-3.el7.x86_64                                       3/6
  Verifying  : mailcap-2.1.41-2.el7.noarch                                  4/6
  Verifying  : apr-util-1.5.2-6.el7.x86_64                                  5/6
  Verifying  : centos-logos-70.0.6-3.el7.centos.noarch                      6/6

Installed:
  httpd.x86_64 0:2.4.6-40.el7.centos

Dependency Installed:
  apr.x86_64 0:1.4.8-3.el7
  apr-util.x86_64 0:1.5.2-6.el7
  centos-logos.noarch 0:70.0.6-3.el7.centos
  httpd-tools.x86_64 0:2.4.6-40.el7.centos
  mailcap.noarch 0:2.1.41-2.el7

Complete!
Loaded plugins: fastestmirror
Cleaning repos: base extras updates
Cleaning up everything
Cleaning up list of fastest mirrors
 ---> 3f9981b5b344
Removing intermediate container 874b59f7d611
Step 6 : RUN echo "Hello Apache." > /var/www/html/index.html
 ---> Running in 35ac84af5292
 ---> af126fbeaa9c
Removing intermediate container 35ac84af5292
Step 7 : RUN systemctl enable httpd
 ---> Running in 1df2c5c8904f
Created symlink from /etc/systemd/system/multi-user.target.wants/httpd.service to /usr/lib/systemd/system/httpd.service.
 ---> 1e3a81ca9e1a
Removing intermediate container 1df2c5c8904f
Step 8 : EXPOSE 80
 ---> Running in 0fb9300324e1
 ---> 003626f91226
Removing intermediate container 0fb9300324e1
Successfully built 003626f91226

```

## おわりに

[Can't install httpd on fedora (cpio: cap_set_file), fixed somewhere or not? · Issue #8966 · docker/docker](https://github.com/docker/docker/issues/8966)
元々は aufs のバグ（？）のようなので、`CONFIG_AUFS_XATTR`が設定できるバージョンまで aufs を更新して、`CONFIG_AUFS_XATTR`を enable にすれば、boot2docker に限らず他の LinuxOS でもこの問題は解消するとのこと。

> thaJeztah commented on 19 Apr 2015
>
> > Basically, this is a duplicate of #6980. This is an issue with (older versions of) AUFS, see #6980 (comment):
> >
> > You need a newer version of AUFS with support for CONFIG_AUFS_XATTR, and that kernel config option must be enabled. That fixes the cpio: cap_set_file on boot2docker (boot2docker/boot2docker#818), and should fix it on any Linux using AUFS.
> > I just tested this issue and was able to reproduce it on a DigitalOcean droplet using AUFS. Switching to BTRFS as storage driver resolved the problem.
> >
> > I also spotted an external issue being linked to this, and there seems to be a 'workaround' here
> > owncloud/core#12967 (comment) which I haven't tested.
> >
> > This is not a bug in Docker, but an issue with older AUFS versions missing CONFIG_AUFS_XATTR support. To resolve the issue, update the AUFS version and enable the CONFIG_AUFS_XATTR attribute, or switch to another storage driver.
> >
> > I'm going to close this issue, because this is a duplicate of #6980 and (unfortunately) nothing Docker itself can resolve.
> >
> > Feel free to continue the discussion, perhaps someone is able to add additional tips.

aufs のバージョンアップするだけの元気が残ってないので、aufs のバージョンを確認する方法を残して終わりにする。

```
[kaakaa@kaakaa-System-Product-Name] ~/work/docker/book-practical-guide/ch5-1-1
% dmesg | fgrep -i aufs
[    5.851021] aufs 3.13-20140303
[    6.457590] aufs au_opts_parse:1155:docker[2298]: unknown option dirperm1
[ 8083.921023] Modules linked in: autofs4 nls_iso8859_1 nls_utf8 udf crc_itu_t veth xt_addrtype aufs pci_stub vboxpci(OX) vboxnetadp(OX) vboxnetflt(OX) ipt_MASQUERADE iptable_nat nf_nat_ipv4 nf_nat nf_conntrack_ipv4 nf_defrag_ipv4 xt_conntrack nf_conntrack ipt_REJECT xt_CHECKSUM iptable_mangle xt_tcpudp bridge stp llc vboxdrv(OX) ip6table_filter ip6_tables iptable_filter ip_tables ebtable_nat ebtables x_tables snd_hda_codec_hdmi snd_hda_codec_via joydev usb_storage hid_generic bnep rfcomm bluetooth hid_logitech_dj usbhid hid intel_rapl snd_hda_intel x86_pkg_temp_thermal snd_hda_codec intel_powerclamp eeepc_wmi asus_wmi snd_hwdep coretemp snd_pcm kvm_intel sparse_keymap kvm snd_page_alloc i915 snd_seq_midi snd_seq_midi_event snd_rawmidi crct10dif_pclmul crc32_pclmul binfmt_misc snd_seq snd_seq_device aesni_intel aes_x86_64 drm_kms_helper lrw gf128mul glue_helper snd_timer ablk_helper drm cryptd snd mei_me i2c_algo_bit mei soundcore shpchp lpc_ich wmi mac_hid video serio_raw parport_pc ppdev lp parport psmouse alx pata_acpi mdio
[107238.762713] aufs au_opts_parse:1155:docker[27263]: unknown option dirperm1
```

参考： [AuFS (Advanced multi layered unification filesystem) のインストール - 私の二次記憶](http://d.hatena.ne.jp/ayokoyama/20140604/p2)
