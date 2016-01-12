# Lastore Daemon

**Description**

Lastore Daemon is based on dbus and support apt backend (And Currently only support apt).

The project is mainly support Linux Application Store. Currently it power deepin store 4.0.

## Dependencies

### Build dependencies
- golang >= 1.2
- pkg-config
- make
- glib-2.0

### Runtime dependencies
- apt-get
- apt-cache
- apt

## Installation

### Debian 8.0 (jessie)

Install prerequisites
```
$ sudo apt-get install make \
                       golang-go \
                       pkg-config \
                       libglib2.0-dev
```

Build
```
$ make
```

```
$ sudo make install
```

Or, generate package files and install Deepin Terminal with it
```
$ debuild -uc -us ...
$ sudo dpkg -i ../lastore-daemon*deb
```

## Usage

### lastore-daemon and lastore-session-helper

lastore-daemon need run as root user. It will autostart by systemd
as system service.

And the lastore-session-helper need run as current session owner user.
It will autostart by dbus-daemon as session service.

There has two group of interface.
The Manager and the Updater. see [Hacking guide] for Detail information.

Normal you don't need use this. Just run deepin-appstore.

But you can use it by tools like d-feet, dbus-send or busctl.

For example, use the PackageDesktopPath api to query the desktop path of
any installed package.

```
dbus-send --print-reply --system --dest=com.deepin.lastore /com/deepin/lastore com.deepin.lastore.Manager.PackageDesktopPath string:"gedit"
```

### lastore-tools
lastore-tools is used generate some index file in /var/lib/lastore
```
 % lastore-tools -h
Usage of ./bin/lastore-tools:
  -item string
    	categories|applications|xcategories|desktop|lastore-remove|lastore-install|update_infos|mirrors
  -output string
    	the file to write
```
There has two group information.

The first group items, update data from network. lastore-daemon will use this to update meta data.
- categories
- applications
- xcategories
- mirrors

And the second group items, update data when local system changed. dpkg hook will use this to update meta data.
- desktop
- update_infos

*NOTE*: Don't use lastore-remove and lastore-install items. The is just for internal testing .
It will install or remove  *ALL OF APPLICATIONS* in dstore, So it very likely to broke your system.


### lastore-smartmirror
It's the helper utils for apt with smartmirror patch. Can't be used alone.


## Getting help

Any usage issues can ask for help via

* [Gitter](https://gitter.im/orgs/linuxdeepin/rooms)
* [IRC channel](https://webchat.freenode.net/?channels=deepin)
* [Forum](https://bbs.deepin.org)
* [WiKi](http://wiki.deepin.org/)

## Getting involved

We encourage you to report issues and contribute changes

* [Contribution guide for users](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Users)
* [Contribution guide for developers](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Developers).
* [Hacking guide for developers](hacking.org)

## License

Lastore Daemon is licensed under [GPLv3](LICENSE).