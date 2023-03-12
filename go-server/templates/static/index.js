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

function load() {
    fetch("/api", {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
        }},)
    .then((response) => response.json())
    .then((data) => {
        // console.log(data)
        document.getElementById("content").innerHTML = jsonToTable(data);
    })
};