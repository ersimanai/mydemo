

var socket;

$(document).ready(function () {
      alert($('#uname').text())

    
    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());

    socket.onmessage = function (event) {
    	var data = JSON.parse(event.data);

         alert(event.data);
        $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + data.Content+ "</li>");
    
    };
    socket.onopen = function(event) {
        alert("连接成功")
    };
    socket.onclose = function(event) {
    	alert("关闭连接")
    }
    // Send messages.
    var postConecnt = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(uname+':'+content);
        $('#sendbox').val("");

    }

    $('#sendbtn').click(function () {
        //alert("send_btn")
        postConecnt();
    });
});