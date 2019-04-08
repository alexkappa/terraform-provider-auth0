#!/bin/bash

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
go.googlesource.com,FALSE,/,TRUE,2147483647,o,git-alex.kalyvitis.gmail.com=1/5iP-Xsg_VUM_AYaQgEiZDcY3Lo4n-OjATOZGEZJlZv62OojU8zybBBYm4LINM4EV
__END__
