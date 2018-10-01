function setup () {
  frameRate(1)
  noCursor()
  createCanvas(windowWidth, windowHeight)
  initSponsors()
  sp = new SponsorPanel()
  lp = new LogoPanel()
}

function draw () {
  update()
  background(255)
  if (sp.isReady()) {
    sp.render()
  }
  lp.render()
}

function windowResized () {
  resizeCanvas(windowWidth, windowHeight)
  sp.setLeftPos(createVector(10, windowHeight - 310))
  sp.setRightPos(createVector(windowWidth - 310, windowHeight - 310))
}

function update () {
  sp.tick()
}
