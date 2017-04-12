
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

var colorBootcamp = "#EF6C78";
var colorSoccer = "#FFFFEE"
var colorFitness = "#DADDF2"
var colorRunning = "#AABBDD"
var colorSwimming = "#8888AA"
var colorFighting = "#3F487F"
var colorCycling = "#223399"
var colorGymnastics = "#56141A"
var colorYoga = "#95222D"
var colorHockey = "#D53140"
var colorDancing = "#FBDC65"

var dayFilterVal = "0";
var dayPartFilterVal = "0";
var dataFilterVal = "0";
var pointOrHood = "point";

var geojson;
var geojsonPoly;

$(document).ready(function () {
    $('#dayfilter').on('change', function (event) {
        if ($("#dataRepresentationSelect")[0].value === "1") {
            dayFilterVal = event.target.value;
            filter();
        }
    })

    $('#daypartfilter').on('change', function (event) {
        if ($("#dataRepresentationSelect")[0].value === "1") {
            dayPartFilterVal = event.target.value;
            filter();
        }
    })

    $('#dataRepresentationSelect').on('change', function (event) {
        dataFilterVal = event.target.value;
        filterData();
    })
});

function filterData() {
    if (dataFilterVal === "1") {
        map.removeSource("polygon");
        map.removeLayer("polygon");
        map.addSource('point', {
            "type": "geojson",
            "data": geojson
        });
        map.addLayer({
            "id": "point",
            "type": "circle",
            "source": "point",
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
    } else {
        map.removeSource("point");
        map.removeLayer("point");
        map.addSource('polygon', {
            "type": "geojson",
            "data": geojsonPoly
        });
        map.addLayer({
            "id": "polygon",
            "type": "fill",
            "source": "polygon",
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

function filter() {
    if (dayFilterVal == "0" && dayPartFilterVal == "0") {
        setDefaultLayer();
        return;
    }
    var resJson = geojson;
    if (dayFilterVal != "0") {
        resJson = filterOnDay(parseInt(dayFilterVal));
    }
    if (dayPartFilterVal != "0") {
        resJson = filterOnDayPart(resJson, parseInt(dayPartFilterVal));
    }

    map.removeSource("point");
    map.removeLayer("point");
    map.addSource('point', {
        "type": "geojson",
        "data": resJson
    });

    map.addLayer({
        "id": "point",
        "type": "circle",
        "source": "point",
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



function filterOnDay(number) {
    if (number == 0) {
        setDefaultLayer();
        return geojson;
    }
    var day = getDayFromNumber(number);

    var resJson = { features: [], type: "FeatureCollection" };
    for (var i = 0; i < geojson.features.length; i++) {
        if (geojson.features[i].properties.day == day) {
            resJson.features.push(geojson.features[i]);
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
        setDefaultLayer();
        return geojson;
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


function setDefaultLayer() {
    map.removeSource("point")
    map.removeLayer("point");
    map.addSource('point', {
        "type": "geojson",
        "data": geojson
    });

    map.addLayer({
        "id": "point",
        "type": "circle",
        "source": "point",
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
            geojsonPoly = geojson;
            for (var cluster of msg.clusters) {
                for (var id of cluster.ids) {
                    for (var f of geojsonPoly.features) {
                        if (f.properties.field_1 == id) {
                            f.properties.category = cluster.category;
                        }
                    }
                }
            }

            map.addSource('polygon', {
                "type": "geojson",
                "data": geojsonPoly
            });

            map.addLayer({
                "id": "polygon",
                "type": "fill",
                "source": "polygon",
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

            var requestPolygons = $.ajax({
                url: "http://localhost:8080/getpolygonsstrava",
                method: "GET",
                dataType: "json"
            });
            requestPolygons.done(function (msg) {
                console.log(msg);
                for(var cluster of msg.clusters) {
                    for(var id of cluster.ids) {
                        for(var f of geojson.features) {
                            if(f.properties.field_1 == id) {
                                f.properties.category = cluster.category;
                                // TODO: SET CATEGORY VAN CLUSTER, DEZE STAAT NU NIET IN DB
                            }
                        }
                    }
                }

                // map.removeSource("point");
                // map.removeLayer("point");
                map.removeSource("polygon");
                map.removeLayer("polygon");
                map.addSource('polygon', {
                    "type": "geojson",
                    "data": geojson
                });
                map.addLayer({
                    "id": "polygon",
                    "type": "fill",
                    "source": "polygon",
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

            });
            requestPolygons.fail(function (jqXHR, textStatus) {
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
        geojson = msg;
    });
    requestPolygons.fail(function (jqXHR, textStatus) {
        alert(textStatus);
    });
}
$(document).ready(function () {
    getPolygons();
    getPoints();
    getStravaPolygons();
});