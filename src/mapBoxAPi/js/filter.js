
// var colorBootcamp = "#042037";
// var colorSoccer = "#496D89"
// var colorFitness = "#114800"
// var colorRunning = "#A1D890"
// var colorSwimming = "#554600"
// var colorFighting = "#D4C26A"
// var colorCycling = "#551800"
// var colorGymnastics = "#D4886A"
// var colorYoga = "#95222D"
// var colorHockey = "#D53140"
// var colorDancing = "#FBDC65"

//
// 
// 
// 
// 
// 
// 
// 
// 

var colorSoccer = "#FFFFEE"
var colorFitness = "#AABBDD"
var colorRunning = "#223399"
var colorSwimming = "#085F12"
var colorFighting = "#51DD61"
var colorCycling = "#FBDC65"
var colorGymnastics = "#886B00"
var colorYoga = "#EF6C78"
var colorHockey = "#D53140"
var colorBootcamp = "#95007F";
var colorDancing = "#56141A"

var dayFilterVal = "0";
var dayPartFilterVal = "0";
var dataFilterVal = "0";
var pointOrHood = "point";

var twitterPointsGeoJson;
var twitterPolygonsGeoJson;
var stravaPolygonsGeoJson;
var stravaPointsGeoJson;

$(document).ready(function () {
    $('#dayfilter').on('change', function (event) {
        if ($("#dataRepresentationSelect")[0].value === "2") {
            dayFilterVal = event.target.value;
            filter(twitterPointsGeoJson);
        } else if ($("#dataRepresentationSelect")[0].value === "3") {
            dayFilterVal = event.target.value;
            filter(stravaPointsGeoJson);
        }
    })

    $('#daypartfilter').on('change', function (event) {
        if ($("#dataRepresentationSelect")[0].value === "2") {
            dayPartFilterVal = event.target.value;
            filter(twitterPointsGeoJson);
        } else if ($("#dataRepresentationSelect")[0].value === "3") {
            dayPartFilterVal = event.target.value;
            filter(stravaPointsGeoJson);
        }
    })

    $('#dataRepresentationSelect').on('change', function (event) {

        dataFilterVal = event.target.value;
        filterData();
    })

    getPolygons();
    getPoints();
    getStravaPoints();
});

function filterData() {
    map.removeSource("datalayer");
    map.removeLayer("datalayer");
    if (dataFilterVal === "2") {

        map.addSource('datalayer', {
            "type": "geojson",
            "data": twitterPointsGeoJson
        });
        map.addLayer({
            "id": "datalayer",
            "type": "circle",
            "source": "datalayer",
            'paint': {
                // make circles larger as the user zooms from z12 to z22
                'circle-radius': {
                    'base': 5,
                    'stops': [[12, 6], [22, 180]]
                },
                // color circles by ethnicity, using data-driven styles
                'circle-color': {
                    property: 'sport-category',
                    type: 'categorical',
                    stops: [
                        ['soccer', colorSoccer],
                        ['fitness', colorFitness],
                        ['running', colorRunning],
                        ['swimming', colorSwimming],
                        ['fightingsport', colorFighting],
                        ['cycling', colorCycling],
                        ['gymnastics', colorGymnastics],
                        ['yoga', colorYoga],
                        ['hockey', colorHockey],
                        ['bootcamp', colorBootcamp],
                        ['dancing', colorDancing]
                    ]
                }
            }
        });
        dayFilterVal = $('#dayfilter')[0].value;
        dayPartFilterVal = $('#daypartfilter')[0].value;
        
        filter(twitterPointsGeoJson);
    } else if (dataFilterVal == "3") {
        map.addSource('datalayer', {
            "type": "geojson",
            "data": stravaPointsGeoJson
        });
        map.addLayer({
            "id": "datalayer",
            "type": "circle",
            "source": "datalayer",
            'paint': {
                // make circles larger as the user zooms from z12 to z22
                'circle-radius': {
                    'base': 5,
                    'stops': [[12, 6], [22, 180]]
                },
                // color circles by ethnicity, using data-driven styles
                'circle-color': {
                    property: 'sport-category',
                    type: 'categorical',
                    stops: [
                        ['soccer', colorSoccer],
                        ['fitness', colorFitness],
                        ['running', colorRunning],
                        ['swimming', colorSwimming],
                        ['fightingsport', colorFighting],
                        ['cycling', colorCycling],
                        ['gymnastics', colorGymnastics],
                        ['yoga', colorYoga],
                        ['hockey', colorHockey],
                        ['bootcamp', colorBootcamp],
                        ['dancing', colorDancing]
                    ]
                }
            }
        });
        dayFilterVal = $('#dayfilter')[0].value;
        dayPartFilterVal = $('#daypartfilter')[0].value;
        filter(stravaPointsGeoJson);
    } else if (dataFilterVal == "1") {
        map.addSource('datalayer', {
            "type": "geojson",
            "data": stravaPolygonsGeoJson
        });

        map.addLayer({
            "id": "datalayer",
            "type": "fill",
            "source": "datalayer",
            'layout': {},
            'paint': {
                'fill-color': {
                    property: 'category',
                    type: 'categorical',
                    stops: [
                        ['soccer', colorSoccer],
                        ['fitness', colorFitness],
                        ['running', colorRunning],
                        ['swimming', colorSwimming],
                        ['fightingsport', colorFighting],
                        ['cycling', colorCycling],
                        ['gymnastics', colorGymnastics],
                        ['yoga', colorYoga],
                        ['hockey', colorHockey],
                        ['bootcamp', colorBootcamp],
                        ['dancing', colorDancing]
                    ]
                },
                'fill-opacity': 0.5
            }
        });
    } else {

        map.addSource('datalayer', {
            "type": "geojson",
            "data": twitterPolygonsGeoJson
        });

        map.addLayer({
            "id": "datalayer",
            "type": "fill",
            "source": "datalayer",
            'layout': {},
            'paint': {
                'fill-color': {
                    property: 'category',
                    type: 'categorical',
                    stops: [
                        ['soccer', colorSoccer],
                        ['fitness', colorFitness],
                        ['running', colorRunning],
                        ['swimming', colorSwimming],
                        ['fightingsport', colorFighting],
                        ['cycling', colorCycling],
                        ['gymnastics', colorGymnastics],
                        ['yoga', colorYoga],
                        ['hockey', colorHockey],
                        ['bootcamp', colorBootcamp],
                        ['dancing', colorDancing]
                    ]
                },
                'fill-opacity': 0.5
            }
        });
    }
}

function filter(currentGeo) {
    if (dayFilterVal == "0" && dayPartFilterVal == "0") {
        setDefaultLayer(currentGeo);
        return;
    }
    var resJson = currentGeo;
    if (dayFilterVal != "0") {
        resJson = filterOnDay(parseInt(dayFilterVal), currentGeo);
    }
    if (dayPartFilterVal != "0") {
        resJson = filterOnDayPart(resJson, parseInt(dayPartFilterVal));
    }
    map.removeSource("datalayer");
    map.removeLayer("datalayer");
    map.addSource('datalayer', {
        "type": "geojson",
        "data": resJson
    });

    map.addLayer({
        "id": "datalayer",
        "type": "circle",
        "source": "datalayer",
        'paint': {
            // make circles larger as the user zooms from z12 to z22
            'circle-radius': {
                'base': 5,
                'stops': [[12, 6], [22, 180]]
            },
            // color circles by ethnicity, using data-driven styles
            'circle-color': {
                property: 'sport-category',
                type: 'categorical',
                stops: [
                    ['soccer', colorSoccer],
                    ['fitness', colorFitness],
                    ['running', colorRunning],
                    ['swimming', colorSwimming],
                    ['fightingsport', colorFighting],
                    ['cycling', colorCycling],
                    ['gymnastics', colorGymnastics],
                    ['yoga', colorYoga],
                    ['hockey', colorHockey],
                    ['bootcamp', colorBootcamp],
                    ['dancing', colorDancing]
                ]
            }
        }
    });
}



function filterOnDay(number, currentGeo) {
    if (number == 0) {
        setDefaultLayer(currentGeo);
        return currentGeo;
    }
    var day = getDayFromNumber(number);

    var resJson = { features: [], type: "FeatureCollection" };
    for (var i = 0; i < currentGeo.features.length; i++) {
        if (currentGeo.features[i].properties.day == day) {
            resJson.features.push(currentGeo.features[i]);
        }
    }
    return resJson;
}

function getDayFromNumber(number) {
    switch (number) {
        case 1: return "Monday";
        case 2: return "Tuesday";
        case 3: return "Wednesday";
        case 4: return "Thursday";
        case 5: return "Friday";
        case 6: return "Saturday";
        case 7: return "Sunday";
        default: return "Monday";
    }
}

function filterOnDayPart(res, number) {
    if (number == 0) {
        setDefaultLayer(res);
        return res;
    }
    var daypart = getDayPartFromNumber(number);

    var resJson = { features: [], type: "FeatureCollection" };
    for (var i = 0; i < res.features.length; i++) {
        if (res.features[i].properties.daypart == daypart) {
            resJson.features.push(res.features[i]);
        }
    }
    return resJson;
}

function getDayPartFromNumber(number) {
    switch (number) {
        case 1: return "Morning";
        case 2: return "Midday";
        case 3: return "Evening";
        case 4: return "Night";
        default: return "Morning";
    }
}


function setDefaultLayer(currentGeo) {
    map.removeSource("datalayer")
    map.removeLayer("datalayer");
    map.addSource('datalayer', {
        "type": "geojson",
        "data": currentGeo
    });

    map.addLayer({
        "id": "datalayer",
        "type": "circle",
        "source": "datalayer",
        'paint': {
            // make circles larger as the user zooms from z12 to z22
            'circle-radius': {
                'base': 5,
                'stops': [[12, 6], [22, 180]]
            },
            // color circles by ethnicity, using data-driven styles
            'circle-color': {
                property: 'sport-category',
                type: 'categorical',
                stops: [
                    ['soccer', colorSoccer],
                    ['fitness', colorFitness],
                    ['running', colorRunning],
                    ['swimming', colorSwimming],
                    ['fightingsport', colorFighting],
                    ['cycling', colorCycling],
                    ['gymnastics', colorGymnastics],
                    ['yoga', colorYoga],
                    ['hockey', colorHockey],
                    ['bootcamp', colorBootcamp],
                    ['dancing', colorDancing]
                ]
            }
        }
    });
}

function getPolygons() {
    loadshp({
        url: 'http://localhost:8080/voronoi', // path or your upload file
        encoding: 'big5', // default utf-8
        EPSG: 3826 // default 4326
    }, function (geojson) {

        var requestPolygons = $.ajax({
            url: "http://localhost:8080/getpolygons",
            method: "GET",
            dataType: "json"
        });
        requestPolygons.done(function (msg) {
            twitterPolygonsGeoJson = geojson;
            for (var cluster of msg.clusters) {
                for (var id of cluster.ids) {
                    for (var f of twitterPolygonsGeoJson.features) {
                        if (f.properties.field_1 == id) {
                            f.properties.category = cluster.category;
                        }
                    }
                }
            }

            getStravaPolygons();

        });
        requestPolygons.fail(function (jqXHR, textStatus) {
            alert(textStatus);
        });
    });
}

function getStravaPolygons() {

    loadshp({
        url: 'http://localhost:8080/voronoistrava', // path or your upload file
        encoding: 'big5', // default utf-8
        EPSG: 3826 // default 4326
    }, function (geojson) {
        stravaPolygonsGeoJson = geojson;
        var requestStravaPolygons = $.ajax({
            url: "http://localhost:8080/getpolygonsstrava",
            method: "GET",
            dataType: "json"
        });
        requestStravaPolygons.done(function (msg) {
            for (var cluster of msg.clusters) {
                for (var id of cluster.ids) {
                    for (var f of stravaPolygonsGeoJson.features) {
                        if (f.properties.field_1 == id) {
                            f.properties.category = cluster.category;
                        }
                    }
                }
            }
            init();
        });
        requestStravaPolygons.fail(function (jqXHR, textStatus) {
            alert(textStatus);
        });
    });

}

function getPoints() {
    var requestPolygons = $.ajax({
        url: "http://localhost:8080/getpoints",
        method: "GET",
        dataType: "json"
    });
    requestPolygons.done(function (msg) {
        twitterPointsGeoJson = msg;
    });
    requestPolygons.fail(function (jqXHR, textStatus) {
        alert(textStatus);
    });
}

function getStravaPoints() {
    var requestPolygons = $.ajax({
        url: "http://localhost:8080/getstravapoints",
        method: "GET",
        dataType: "json"
    });
    requestPolygons.done(function (msg) {
        stravaPointsGeoJson = msg;
    });
    requestPolygons.fail(function (jqXHR, textStatus) {
        alert(textStatus);
    });
}



var loaded = false;
function init() {
    if (!loaded) {
        loaded = true;
        map.addSource('datalayer', {
            "type": "geojson",
            "data": twitterPolygonsGeoJson
        });

        map.addLayer({
            "id": "datalayer",
            "type": "fill",
            "source": "datalayer",
            'layout': {},
            'paint': {
                'fill-color': {
                    property: 'category',
                    type: 'categorical',
                    stops: [
                        ['soccer', colorSoccer],
                        ['fitness', colorFitness],
                        ['running', colorRunning],
                        ['swimming', colorSwimming],
                        ['fightingsport', colorFighting],
                        ['cycling', colorCycling],
                        ['gymnastics', colorGymnastics],
                        ['yoga', colorYoga],
                        ['hockey', colorHockey],
                        ['bootcamp', colorBootcamp],
                        ['dancing', colorDancing]
                    ]
                },
                'fill-opacity': 0.5
            }
        });
    }
}