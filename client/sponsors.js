var sponsorImages = []
var sponsorsLeftTimer = 0
var sponsorsLefti = 0
var sponsorsRightTimer = 225
var sponsorsRighti = 11
var sponsorsRefreshAt = 450
var sponsorsReady = false

function initSponsors () {
  loadJSON('/api/sponsors', function (list) {
    addSponsorImages(list)
  })
}

function addSponsorImages (list) {
  if (list.length !== sponsorImages.length) {
    sponsorsReady = false
    sponsorsImages = []
    for (let i = 0; i < list.length; i++) {
      sponsorImages.push(loadImage('/img/sponsors/' + list[i]))
    }
    sponsorsReady = true
  }
}

function sponsorsRender () {
  if (sponsorImages.length > 5) {
    var l, r
    l = sponsorsLefti
    r = sponsorsRighti
    imageMode(CORNER)
    image(sponsorImages[l], 10, windowHeight - 310)
    image(sponsorImages[r], windowWidth - 310, windowHeight - 310)
  }
}

function sponsorsTick () {
  sponsorsLeftTimer++
  sponsorsRightTimer++

  if (sponsorImages.length > 0 && sponsorsLeftTimer >= sponsorsRefreshAt) {
    sponsorsLeftTimer = 0
    if (sponsorsLefti + 2 >= sponsorImages.length) {
      sponsorsLefti = 0
    } else {
      sponsorsLefti++
      sponsorsLefti++
    }
  }

  if (sponsorImages.length > 0 && sponsorsRightTimer >= sponsorsRefreshAt) {
    sponsorsRightTimer = 0
    if (sponsorsRighti + 2 >= sponsorImages.length) {
      sponsorsRighti = 1
    } else {
      sponsorsRighti++
      sponsorsRighti++
    }
  }
}
