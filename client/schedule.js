/* global
   loadJSON
*/

var presentationList = []
var eventList = []
var scheduleReady = false

class Presentation {
  constructor() {
    this.name
    this.description
    this.location
    this.startTime
    this.endTime
    this.speakers = []
    this.topic
  }
}

class Event {
  constructor() {
    this.name
    this.description
    this.location
    this.startTime
    this.endTime
  }
}

/* exported initSchedule */
function initSchedule() {
  loadJSON('/api/schedule', populateSchedule)
}

function populateSchedule(blob) {
  if (blob.Presentations.length !== presentationList.length) {
    scheduleReady = false
    presentationList = []
    for (let i = 0; i < blob.Presentations.length; i++) {
      let p = new Presentation()
      p.name = blob.Presentations[i].Name
      p.description = blob.Presentations[i].Description
      p.location = blob.Presentations[i].Location
      p.startTime = blob.Presentations[i].StartTime
      p.endTime = blob.Presentations[i].EndTime
      p.speakers = blob.Presentations[i].Speakers
      p.topic = blob.Presentations[i].Topic
      presentationList.push(p)
    }
  }
  if (blob.Events.length !== eventList.length) {
    scheduleReady = false
    eventList = []
    for (let i = 0; i < blob.Events.length; i++) {
      let e = new Event()
      e.name = blob.Events[i].Name
      e.description = blob.Events[i].Description
      e.location = blob.Events[i].Location
      e.startTime = blob.Events[i].StartTime
      e.endTime = blob.Events[i].EndTime
      eventList.push(e)
    }
  }
  scheduleReady = true
}

/* exported SchedulePanel */
class SchedulePanel {
  constructor() {
    this.presentations = []
    this.events = []
    this.panelReady = false
  }

  tick() {
    if (scheduleReady && !this.panelReady) {
      this.presentations = presentationList
      this.events = eventList
      this.panelReady = true
    }
  }
}
