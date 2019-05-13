#!/usr/bin/env bash

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
.googlesource.com,FALSE,/,TRUE,2147483647,o,git-alex.kalyvitis.gmail.com=1/Yw1laTtbLaoiBdXMq9IYwB_eLHgEFA7LRhdCIjqtnrs
__END__