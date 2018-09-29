var clientState = 'INIT'

var imgLogo
var imgWifi
var imgMap

function setup () {
  frameRate(30)
  noCursor()

  createCanvas(windowWidth, windowHeight)

  imgLogo = loadImage('/img/logo-17x.png')
  imgWifi = loadImage('/img/wifi.png')
  // imgMap = loadImage('/img/convention-center-map.png')

  initSponsors()

  sp = new sponsorpanel()

  clientState = 'READY'
}

function draw () {
  update()
  if (clientState === 'READY') {
    // background(223, 202, 150)
    background(255)
    // sponsorsRender()
    sp.render()
    drawHeader()
  }
}

function windowResized () {
  resizeCanvas(windowWidth, windowHeight)
}

function update () {
  // sponsorsTick()
  sp.tick()
}

function drawHeader () {
  fill(255)
  stroke(255)
  // rect(-5, -5, windowWidth, 190)
  imageMode(CENTER)
  image(imgWifi, windowWidth - 200, 100)
  image(imgLogo, 190, 100)
}
