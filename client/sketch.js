/* global sp lp SponsorPanel LogoPanel
   createCanvas initSponsors background resizeCanvas
   windowWidth windowHeight createVector frameRate
   noCursor
*/

/* exported setup */
function setup() {
  frameRate(1)
  noCursor()
  createCanvas(windowWidth, windowHeight)
  initSponsors()

  // .eslintrc: "no-global-assign": ["error", {"exceptions": ["sp", "lp"]}]
  sp = new SponsorPanel()
  lp = new LogoPanel()
}

/* exported draw */
function draw() {
  update()
  background(255)
  if (sp.isReady()) {
    sp.render()
  }
  lp.render()
}

/* exported windowResized */
function windowResized() {
  resizeCanvas(windowWidth, windowHeight)
  sp.setLeftPos(createVector(10, windowHeight - 310))
  sp.setRightPos(createVector(windowWidth - 310, windowHeight - 310))
}

/* exported update */
function update() {
  sp.tick()
}
