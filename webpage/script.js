function Clock() {
  var now = new Date();
  var hour = now.getHours();
  var min  = ('0'+now.getMinutes()).slice(-2);
  //var sec  = ('0'+now.getSeconds()).slice(-2);

  var msg = "現在時刻 " + hour + ":" + min ;
  document.getElementById("TimeArea").innerHTML = msg;
}

setInterval('Clock()',1000);

function CurtainMove() {
  console.log("カーテンの操作が行われた");
}