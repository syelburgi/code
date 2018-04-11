import sys
import string
import math
import re

url_list = ["www.google.com",
            "www.yahoo.com",
            "www.vihal.com",
            "www.sandeep.com"
            ]

char_map = [str(c) for c in range(0, 10)] + [s for s in string.ascii_lowercase] + [s for s in string.ascii_uppercase]


def fetch_the_url(url):
    global url_list

    try:
        index = url_list.index(url)
        return index
    except ValueError:
        url_list.append(url)
        index = url_list.index(url)
        print("url %s newly added" % url)
        return index
    except Exception as e:
        print("fetch the url failed %s", e)
        return -1

def index_2_murl(index):
    murl = ""

    if index == 0:
        return char_map[0]

    while index:
        index, remain = divmod(index, len(char_map))
        murl = char_map[remain] + murl

    return murl

def url_2_miniurl(url):

    if url:
        index = fetch_the_url(url)
        if index != -1:
            print("url %s found at %d" % (url, index))
            murl = index_2_murl(index)
            print("murl %s" %murl)
            return 0, "https://" + murl
        else:
            return -1, ""
    else:
        print("invalid url")
        return -1, ""

def miniurl_2_url(miniurl):

    try:
        domain = re.search('https://(.+?)', miniurl).group(1)
    except AttributeError:
        print("miniurl should be of format https://xyz")
        return -1, ""

    try:
        i = 0
        val = 0
        for c in domain:
            index = char_map.index(c)
            val += index * math.pow(len(char_map), i)
            i += 1

        return 0, url_list[int(val)]
    except ValueError:
        print("invalid url")
    except Exception as e:
        print("exception occured %s", str(e))
    return -1, ""

if __name__ == "__main__":


    print(sys.argv[0])
    print(sys.argv[1])
    print(sys.argv[2])
    print(char_map)

    if sys.argv[1] == "url":
        retv, ret = url_2_miniurl(sys.argv[2])

    if sys.argv[1] == "miniurl":
        retv, ret = miniurl_2_url(sys.argv[2])

    print("%s --> %s"%(sys.argv[2], ret))