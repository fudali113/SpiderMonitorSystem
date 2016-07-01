<!DOCTYPE html>

<html>
<head>
  <title>look</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  
</body>
<script src="//cdn.bootcss.com/jquery/3.0.0/jquery.js"></script>
<script type="text/javascript">
	var websocket = new WebSocket("ws://localhost:8080/ws");
	websocket.onopen = function(evt) { 
            alert('websocket连接成功') 
        }; 
        websocket.onclose = function(evt) { 
			alert(evt)
            alert('websocket断开连接') 
        }; 
        websocket.onmessage = function(evt) { 
            alert(evt.data) 
        }; 
        websocket.onerror = function(evt) { 
			alert(evt)
            alert('websocket出现错误') 
        }; 
</script>
</html>
