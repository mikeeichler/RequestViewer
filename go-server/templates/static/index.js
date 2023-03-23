function twelve13() {
    const isIPhone12 = window.innerWidth === 390 && window.innerHeight === 844 && window.devicePixelRatio === 3;
    const isIPhone13 = window.innerWidth === 390 && window.innerHeight === 844 && window.devicePixelRatio === 3.125;

    if (isIPhone12) {
        return "This is an iPhone 12";
    } else if (isIPhone13) {
        return "This is an iPhone 13";
    } else {
        return "This is not an iPhone 12 or iPhone 13";
    }
}

function jsonToTable(j) {
    let table = "<table><tr><th>Header</th><th>Value</th></tr>";
  for (const k in j) {
    table += `<tr><td>${k}</td><td>${j[k]}</td></tr>`;
    if (k == "timestamp") {
      timestamp = j[k];
    }
  }
  table += '</table>'
  return table;
}

function load(device) {
  // console.log(DeviceAtlas.)
      const videoObject = new Map();
      var video = document.createElement('video');
      if (video.canPlayType('video/mp4; codecs="ap4h.2, mp4a.40.2"')) {
          videoObject.set('PR422', true)
      } else {
        videoObject.set('PR422', false)
      }
      videoObject.set('innerHeight', window.innerHeight)
      videoObject.set('screenHeight', screen.height)
      videoObject.set('screenAvailHeight', screen.availHeight)
      videoObject.set('devicePixelRatio', window.devicePixelRatio)
        
    fetch("/api", {
        method: 'GET',
        headers: {
            DAPROPS,
            'Accept': 'application/json',
            'X-PR422': videoObject.get('PR422'),
            'X-PR4444': videoObject.get('PR4444'),
            'X-WEBGL-HASH': videoObject.get('webglHash'),
            'X-ACTUAL-DEVICE': device,
            'X-INNER-HEIGHT': videoObject.get('innerHeight'),
            'X-SCREEN-HEIGHT': videoObject.get('screenHeight'),
            'X-SCREEN-AVAIL-HEIGHT': videoObject.get('screenAvailHeight'),
            'X-DPR': videoObject.get('devicePixelRatio'),
            'X-MESSAGE': document.getElementById("message").value
        }},)
    .then((response) => response.json())
    .then((data) => {
        console.log(data)
        document.getElementById("message").value = ''
        // window.location.reload(true)
        document.getElementById("content").style.visibility = "visible";
        document.getElementById("content").innerHTML = jsonToTable(data);
        // document.write(jsonToTable(data))
    })
    const tw13 = twelve13();
    console.log(tw13)
};