# go-dns

go-dns is a dynamic IP updater for DNSimple.com domains.  go-dns will update
and pre-configured domain's ip to the ip of your router.

## Quick Start

Download one of the archives listed below, and extract to the location of
your choice.  One of the included files is a settings.json.sample.  Rename
that file to settings.json, and update its contents appropriately.  The
values in the file are pretty self explanitory.  The required email is
the one you use to log into dnsimple.com with.  The token can be found
in your dnsimple.com account profile, under the API tab.

Here is an example `settings.json` file.  Note, you shouldn't have to change
the record-type attribute.
```json
{
  "credentials" : {
      "email": "john.doe@example.com",
      "token": "ABCfGHI56aPe98IuYTR"
  },
  "domains" : [
    {
      "name": "example.com",
      "record-type": "A"
    },
    {
      "name": "foobarbaz.com",
      "record-type": "A"
    }
  ]
}
```

Next, run go-dns

```
$ go-dns
...
```

And thats it.  go-dns will print out your router ip, and if the IPs of
the domains match or not.  If go-dns finds a mismatch, it will update the
domain IP to the router's IP.

## Available Downloads

* Mac OS X [amd64](http://dl.bintray.com/jcarley/go-dns/0.1_darwin_amd64.zip) | [386](http://dl.bintray.com/jcarley/go-dns/0.1_darwin_386.zip)
* Windows [amd64](http://dl.bintray.com/jcarley/go-dns/0.1_windows_amd64.zip) | [386](http://dl.bintray.com/jcarley/go-dns/0.1_windows_386.zip)
* Linux [amd64](http://dl.bintray.com/jcarley/go-dns/0.1_linux_amd64.zip) | [386](http://dl.bintray.com/jcarley/go-dns/0.1_linux_386.zip) | [arm](http://dl.bintray.com/jcarley/go-dns/0.1_linux_arm.zip)


## Contributing to go-dns

* Fork this repository on github
* Make your changes and send us a pull request
* If we like them we'll merge them