class sponsorpanel {
  constructor () {
    this.imgList = []
    this.leftPos = createVector(10, windowHeight - 310)
    this.rightPos = createVector(windowWidth - 310, windowHeight - 310)
    this.hidePos = createVector(-600, -600)
    this.leftIndex = int(random(0, this.imgList.length))
    this.rightIndex = int(random(0, this.imgList.length))
    this.maxTimer = 450
    this.timer = this.maxTimer
    this.panelReady = false
  }

  tick () {
    if (sponsorsReady === true) {
      if (this.imgList.length !== sponsorImages.length) {
        this.imgList = []
        for (let i = 0; i < sponsorImages.length; i++) {
          this.imgList.push(new img(this.hidePos, sponsorImages[i]))
        }
        this.panelReady = true
      }
      this.timer++
      this.leftPos = createVector(10, windowHeight - 310)
      this.rightPos = createVector(windowWidth - 310, windowHeight - 310)
      if (this.timer >= this.maxTimer) {
        this.timer = 0
        this.imgList[this.leftIndex].setPosVec(this.hidePos)
        this.leftIndex = int(random(0, this.imgList.length))
        this.imgList[this.leftIndex].setPosVec(this.leftPos)
      }
      if (this.timer === this.maxTimer / 2) {
        this.imgList[this.rightIndex].setPosVec(this.hidePos)
        this.rightIndex = int(random(0, this.imgList.length))
        this.imgList[this.rightIndex].setPosVec(this.rightPos)
      }
    }
  }

  render () {
    if (sponsorsReady === true && this.panelReady === true) {
      this.imgList[this.leftIndex].render()
      this.imgList[this.rightIndex].render()
    }
  }
}
