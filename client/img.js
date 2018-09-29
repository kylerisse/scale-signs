class img {
  constructor (posVec, image) {
    this.pos = posVec
    this.vel = createVector(0, 0)
    this.dest = posVec
    this.image = image
    this.maxspeed = 4
    this.maxforce = 0.1
  }

  tick () {
    var desired = p5.Vector.sub(this.dest, this.pos)
    var d = desired.mag()
    if (d < 100) {
      var m = map(d, 0, 100, 0, this.maxspeed)
      desired.setMag(m)
    } else {
      desired.setMag(this.maxspeed)
    }

    var steer = p5.Vector.sub(desired, this.vel)
    steer.limit(this.maxforce)
    this.applyForce(steer)
  }

  render () {
    imageMode(CORNER)
    image(this.image, this.pos.x, this.pos.y)
  }

  setDest (x, y) {
    this.dest = createVector(x, y)
  }

  setVel (x, y) {
    this.vel = createVector(x, y)
  }

  setPos (x, y) {
    this.pos = createVector(x, y)
  }

  setPosVec (v) {
    this.pos = v
  }
}
