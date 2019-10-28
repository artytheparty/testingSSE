function OnLoaded(){
    var source = new EventSource("/sse/dashboard");
    source.onopen = function (event) {
        console.log("OnMessage called:");
        console.dir(event);
        document.getElementById("counter").innerHTML= event.data;
    }
}