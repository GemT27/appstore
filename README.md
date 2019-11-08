# appstore
Golang to scrape application data from the iTunes App Store.

## Installation
```
go get github.com/TT527/appstore
```

## Usage
Available methods:
- [app](#app): Query similar software information.
- [keyword](#keyword): Query keyword search results.
- [iTunes](#iTunes): Search results for iTunes.

[About Country Code](https://affiliate.itunes.apple.com/resources/documentation/linking-to-the-itunes-music-store/#appendix)
### app
 Query similar software information:

* `term`: the term to search.
* `country`: the two letter country code to get the app details from. Note this also affects the language of the data.
* `num`: the amount of elements to retrieve.

Example:

```javascript
result,_:=search.App2Json("google", "US",2)
fmt.Println(result)
```

Results:

```javascript
{
    "resultCount": 2,
    "results": [
        {
            "artistViewUrl": "https://apps.apple.com/us/developer/google-llc/id281956209?uo=4",
            "artworkUrl60": "https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/ae/0b/95/ae0b95c1-e86a-7710-683d-c444a08cfd18/source/60x60bb.jpg",
            "artworkUrl100": "https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/ae/0b/95/ae0b95c1-e86a-7710-683d-c444a08cfd18/source/100x100bb.jpg",
            "screenshotUrls": [
                "https://is2-ssl.mzstatic.com/image/thumb/Purple123/v4/cc/95/b4/cc95b4c6-5275-0053-5f9c-e369c97fc0e5/pr_source.png/696x696bb.png",
                "https://is5-ssl.mzstatic.com/image/thumb/Purple123/v4/ec/20/9b/ec209b97-5041-00c2-1f40-ab6a6c0f4a60/pr_source.png/696x696bb.png",
                "https://is2-ssl.mzstatic.com/image/thumb/Purple123/v4/9b/d7/ba/9bd7bae0-211b-66f3-0a37-603dc952b30e/pr_source.png/696x696bb.png",
                "https://is1-ssl.mzstatic.com/image/thumb/Purple123/v4/59/5b/cc/595bcccc-deab-bb2b-4146-a4d08a63e927/pr_source.png/696x696bb.png",
                "https://is3-ssl.mzstatic.com/image/thumb/Purple113/v4/ab/92/fd/ab92fd2b-65fd-2f9b-586e-b0600ac3a0ec/pr_source.png/696x696bb.png"
            ],
            "ipadScreenshotUrls": [
                "https://is4-ssl.mzstatic.com/image/thumb/Purple123/v4/94/5d/b1/945db107-7cd4-eba6-1fda-eaa778657a96/mzl.mttavtfx.jpg/552x414bb.jpg",
                "https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/bb/53/05/bb530509-2d1c-d427-62bc-6ee3cddd77f5/mzl.pwucdfyy.jpg/552x414bb.jpg",
                "https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/c8/19/7c/c8197c49-139c-21ff-d919-a95549dcf96b/mzl.jhguvtbs.png/552x414bb.png",
                "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/61/48/fc/6148fc12-ee6e-2041-5047-5f2429fe5d0b/mzl.lgufspcq.jpg/552x414bb.jpg"
            ],
            "appletvScreenshotUrls": [],
            "artworkUrl512": "https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/ae/0b/95/ae0b95c1-e86a-7710-683d-c444a08cfd18/source/512x512bb.jpg",
            "advisories": [
                "Unrestricted Web Access"
            ],
            "supportedDevices": [
                "iPhone5s-iPhone5s",
                ...
            ],
            "isGameCenterEnabled": false,
            "kind": "software",
            "features": [
                "iosUniversal"
            ],
            "trackCensoredName": "Google",
            "trackViewUrl": "https://apps.apple.com/us/app/google/id284815942?uo=4",
            "contentAdvisoryRating": "17+",
            "fileSizeBytes": "249401344",
            "averageUserRatingForCurrentVersion": 4,
            "languageCodesISO2A": [
                "AR",
                "CA",
                ...
            ],
            "sellerUrl": "http://www.google.com/search/about",
            "userRatingCountForCurrentVersion": 317,
            "trackContentRating": "17+",
            "currentVersionReleaseDate": "2019-07-16T22:26:00Z",
            "releaseNotes": "• Bug fixes and performance improvements\n \nWe are always working to make the app faster and more stable. If you are enjoying the app, please consider leaving a review or rating!",
            "releaseDate": "2008-07-11T07:00:00Z",
            "trackId": 284815942,
            "trackName": "Google",
            "sellerName": "Google LLC",
            "primaryGenreId": 6002,
            "primaryGenreName": "Utilities",
            "genreIds": [
                "6002",
                "6006"
            ],
            "currency": "USD",
            "wrapperType": "software",
            "version": "78.0",
            "artistId": 281956209,
            "artistName": "Google LLC",
            "genres": [
                "Utilities",
                "Reference"
            ],
            "price": 0,
            "description": "The Google app keeps you in the know about things that matter to you. Find quick answers, explore your interests, and stay up to date with Discover. The more you use the Google app, the better it gets.\n\nSearch and browse:\n• Nearby shops and restaurants\n• Live sports scores and schedules\n• Movies times, casts, and reviews\n• Videos and images\n• News, stock information, and more\n• Anything you’d find on the web\n\nGet personalized updates in Discover:\n• Stay in the know about topics that interest you\n• Start your morning with weather and top news\n• Get real-time updates on sports, movies, and events\n• Know as soon as your favorite artists drop new albums\n• Get stories about your interests and hobbies\n• Follow interesting topics, right from Search results \n\nMore ways to access Google:\n• Google Lens — Search what you see with your camera, copy and translate text, find similar apparel, identify plants and animals, scan QR codes and more.\n• iMessage extension — Search and share restaurants, GIFs, and more, without leaving your conversation.\n• Search Google extension — While browsing in Safari, you can share a web page with Google to see suggestions for related content—no need to type anything new in the search box. Tap on the Search Google icon from Safari’s share menu to get started. \n• Gboard — access Google Search, right from your keyboard. Gboard is a keyboard that lets you search and send information, GIFs, emoji, and more—right from your keyboard, in any app. Tap “Gboard” in your app settings to get started. \n• Trending on Google widget — find out what’s trending in your area with our Trending on Google widget. \n\nLearn more about what the Google app can do for you: http://www.google.com/search/about\n\nYour feedback helps us create products you'll love. Join a user research study here:\nhttps://goo.gl/kKQn99",
            "isVppDeviceBasedLicensingEnabled": true,
            "bundleId": "com.google.GoogleMobile",
            "minimumOsVersion": "11.0",
            "formattedPrice": "Free",
            "userRatingCount": 699672,
            "averageUserRating": 4
        },
        {
          ...
        }
    ]
}
```

### keyword
 Query keyword search results:

* `term`: the term to search.
* `country`: the two letter country code to get the app details from. Note this also affects the language of the data.


```javascript
result, _ := search.KeyWord2Json("facebook", "US")
fmt.Println(result)
```

Results:

```javascript
{
    "facebook": "10581",
    "facebook ads manager": "4810",
    "facebook likes": "4718",
    "facebook lite": "5371",
    "facebook marketplace": "6193",
    "facebook messenger": "5768",
    "facebook messenger app": "7682",
    "facebook page": "5184",
    "facebook page manager": "5308",
    "facebook video downloader": "4924"
}
```

### iTunes
  Search results for iTunes:

* `term`: the term to search.
* `country`: the two letter country code to get the app details from. Note this also affects the language of the data.

Example:

```js
result, _ := search.Itunes2Json("facebook", "US")
fmt.Println(result)
```

Returns:

```js
[
    {
        "id": 1,
        "logo": "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/e7/38/00/e7380026-e27f-2e94-9389-c39552b1b02a/source/75x75bb.jpg",
        "name": "Facebook",
        "genre": "Social Networking",
        "href": "https://apps.apple.com/us/app/facebook/id284882215"
    },
    {
        "id": 2,
        "logo": "https://is2-ssl.mzstatic.com/image/thumb/Purple123/v4/5a/e6/9d/5ae69de8-6f03-9287-0628-e35c7acf62a2/source/75x75bb.jpg",
        "name": "Messenger",
        "genre": "Social Networking",
        "href": "https://apps.apple.com/us/app/messenger/id454638411"
    },
    {
        "id": 3,
        "logo": "https://is4-ssl.mzstatic.com/image/thumb/Purple123/v4/a1/a2/42/a1a2429e-d9c5-f207-1a50-6e06349749e7/source/75x75bb.jpg",
        "name": "Instagram",
        "genre": "Photo & Video",
        "href": "https://apps.apple.com/us/app/instagram/id389801252"
    },
    {
        "id": 4,
        "logo": "https://is2-ssl.mzstatic.com/image/thumb/Purple123/v4/f1/39/00/f13900ee-9dbd-86c6-11b1-9c74fb3aff33/source/75x75bb.jpg",
        "name": "Messenger Kids",
        "genre": "Social Networking",
        "href": "https://apps.apple.com/us/app/messenger-kids/id1285713171"
    },
    {
        "id": 5,
        "logo": "https://is4-ssl.mzstatic.com/image/thumb/Purple113/v4/95/ce/a8/95cea82b-a371-3be3-0e14-5c18e41f943c/source/75x75bb.jpg",
        "name": "Workplace by Facebook",
        "genre": "Business",
        "href": "https://apps.apple.com/us/app/workplace-by-facebook/id944921229"
    },
    ...
    ...
    ...
]
```
