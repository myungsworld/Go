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
    });

    var addMessage = function(data) {
        var text = "";
        if (!isBlank(data.name)){
            text = '<strong>' + data.name + ':</strong>'
        }
        text += data.msg
        $chatlog.prepend('<div><span>' + text + '</span></div>')
    }

    addMessage({
        msg : 'hello',
        name: 'aaa'
    })

    addMessage({
        msg : 'hello2'
    })
})