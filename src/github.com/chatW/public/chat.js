$(function(){
    if (!window.EventSource) {
        alert("No EventSource")
        return
    }
    //html에 있는 id 값을 가져옴
    var $chatlog = $('#chat-log')
    var $chatmsg = $('#chat-msg')

    var isBlank = function(string) {
        // == 는 값만 반환 === 는 타입까지 반환
        return string == null || string.trim() === "";
    };

    var username;
    while (isBlank(username)) {
        username = prompt("What's your name?");
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>');
        }
    }

    $('#input-form').on('submit', function(e){
        $.post('/messages', {
            msg: $chatmsg.val(),
            name : username
        })
        $chatmsg.val("")
        $chatmsg.focus()
        //true로 할시 다른페이지로 넘어감
        return false;
    });

    var addMessage = function(data) {
        var text = "";
        if (!isBlank(data.name)){
            text = '<strong>' + data.name + ':</strong>'
        }
        text += data.msg
        $chatlog.prepend('<div><span>' + text + '</span></div>')
    }

    var es = new EventSource('/stream');
    es.onopen = function(e) {
        $.post('/users',{
            name : username
        });
    }
    es.onmessage = function(e) {
       var msg = JSON.parse(e.data);
       addMessage(msg);
    }
    // 윈도우가 닫히기 전에 호출
    window.onbeforeunload = function() {
       $.ajax({
           url: "/users?username=" + username,
           type : "DELETE"
       });
       es.close()
    }
})