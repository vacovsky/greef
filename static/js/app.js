"use strict";

function CtoF(n) {
  return n * (9 / 5) + 32;
}

function getLastValue(lst) {
  return lst[lst.length - 1];
}

function populateTemplate(data) {}

var CHART_OPTIONS = {
  scales: {
    xAxes: [
      {
        type: "time",
        time: {
          unit: "day"
        }
      }
    ]
  }
};

function loadLightingChart() {
  var Http = new XMLHttpRequest();
  var url = "/lightingchartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "Aquarium Lighting",
              borderColor: window.chartColors.orange,
              data: chartData.Data[0],
              fill: false
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },
            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Birghtness"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "% max"
                },
                ticks: {
                  beginAtZero: true,
                  max: 100,
                  min: 0,
                  fontColor: "#bbb",
                  stepSize: 20
                }
              }
            ]
          }
        }
      };
      var ctx = document.getElementById("lighting-chart").getContext("2d");
      new Chart(ctx, config);
      weld(
        document.getElementById("#last-lighting"),
        getLastValue(chartData.Data[0]) + " %"
      );
    }
  };
}

function loadCO2Chart() {
  var Http = new XMLHttpRequest();
  var url = "/co2chartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "Air CO₂ (ppm)",
              borderColor: window.chartColors.green,
              data: chartData.Data[0],
              fill: false
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },
            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Ambient Air CO₂"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "ppm"
                },
                ticks: {
                  beginAtZero: true,
                  min: 0,
                  fontColor: "#bbb",
                  stepSize: 400
                }
              }
            ]
          }
        }
      };
      var ctx = document.getElementById("co2-chart").getContext("2d");
      new Chart(ctx, config);
      weld(
        document.getElementById("#last-air-co2ppm"),
        getLastValue(chartData.Data[0]) + " ppm"
      );
    }
  };
}

function loadTDSChart() {
  var Http = new XMLHttpRequest();
  var url = "/tdschartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "Total Dissolved Solids (TDS PPM)",
              borderColor: window.chartColors.green,
              data: chartData.Data[0],
              fill: false
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },

            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Water Quality"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "ppm"
                },
                ticks: {
                  beginAtZero: true,
                  min: 0,
                  // max: 100000,
                  fontColor: "#bbb"
                  // stepSize: 20000
                }
              }
            ]
          }
        }
      };
      var ctx = document.getElementById("tds-chart").getContext("2d");
      new Chart(ctx, config);
      weld(
        document.getElementById("#last-water-tdsppm"),
        getLastValue(chartData.Data[0]) + " ppm"
      );
    }
  };
}

function loadPHChart() {
  var Http = new XMLHttpRequest();
  var url = "/phchartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "pH",
              borderColor: window.chartColors.green,
              data: chartData.Data[0],
              fill: false
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },

            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Water Acidity"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "pH"
                },
                ticks: {
                  beginAtZero: true,
                  min: 7,
                  max: 9,
                  fontColor: "#bbb",
                  stepSize: 0.2
                }
              }
            ]
          }
        }
      };
      var ctx = document.getElementById("ph-chart").getContext("2d");
      new Chart(ctx, config);
      weld(
        document.getElementById("#last-water-ph"),
        getLastValue(chartData.Data[0])
      );
    }
  };
}

function loadTempChart() {
  var Http = new XMLHttpRequest();
  var url = "/tempchartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      //      populateTemplate();
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "Water Temperature °F",
              borderColor: window.chartColors.blue,
              data: chartData.Data[0],
              fill: false
            },
            {
              label: "Ambient Air Temperature °F",
              fill: false,
              borderColor: window.chartColors.red,
              data: chartData.Data[1]
            },
            {
              label: "Ambient Humidity %",
              fill: false,
              borderColor: window.chartColors.orange,
              data: chartData.Data[2]
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },

            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Environmental"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "°F / %"
                },
                ticks: {
                  beginAtZero: true,
                  min: 0,
                  max: 100,
                  fontColor: "#bbb",
                  stepSize: 20
                }
              }
            ]
          }
        }
      };
      var ctx = document.getElementById("temp-chart").getContext("2d");
      new Chart(ctx, config);

      weld(
        document.getElementById("#last-water-temp"),
        getLastValue(chartData.Data[0]) + " °F"
      );
      weld(
        document.getElementById("#last-air-temp"),
        getLastValue(chartData.Data[1]) + " °F"
      );
      weld(
        document.getElementById("#last-humidity"),
        getLastValue(chartData.Data[2]) + " %"
      );
    }
  };
}

window.chartColors = {
  red: "rgb(255, 99, 132)",
  orange: "rgb(255, 159, 64)",
  yellow: "rgb(255, 205, 86)",
  green: "rgb(75, 192, 192)",
  blue: "rgb(54, 162, 235)",
  purple: "rgb(153, 102, 255)",
  grey: "rgb(201, 203, 207)"
};

function looper() {
  setTimeout(function() {
    loadTempChart();
    loadPHChart();
    loadTDSChart();
    loadCO2Chart();
    loadLightingChart();
    looper();
    console.log("Looping");
  }, 300000);
}

loadTempChart();
loadPHChart();
loadTDSChart();
loadCO2Chart();
loadLightingChart();
looper();
