<!DOCTYPE html>

<html>
<head>
  <title>Spider Monitor System</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.6/css/bootstrap.min.css" />
<link rel="stylesheet" href="/static/css/index.css" />
<body ng-app="monitor">
<script language="javascript" type="text/javascript" src="http://pv.sohu.com/cityjson"></script>
<script> window.ppSettings = {app_uuid:'75cf8791-62d0-11e6-8133-206a8a685554',user_name: returnCitySN.cname+'-用户'};(function(){var w=window,d=document;function l(){var a=d.createElement('script');a.type='text/javascript';a.async=!0;a.src='https://anquanqiao.com/ppcom/assets/pp-library.min.js';var b=d.getElementsByTagName('script')[0];b.parentNode.insertBefore(a,b)}w.attachEvent?w.attachEvent('onload',l):w.addEventListener('load',l,!1);})()</script>
<div id="all" class="panel panel-info" ng-controller="computers" style="padding:15px 15px 0 px 15px">
<img src="/static/img/1.png">
<img src="/static/img/2.png">
<img src="/static/img/3.png">
	<div class="panel-heading">
	  	<div class="panel-title">
					<a href="#">
					<strong>Spider Monitor System</strong>
					</a>
				<div style="float:right;">
				<smss></smss>
				</div>
			</div>
		</div>
	<div class="panel-body">
		<div ng-show="showDetails">
			<setting></setting>
		</div>
		<computer ng-repeat="computer in computers"></computer>
	</div>
	<input class="example1" type="button" value="button">
</div>
</body>
<script src="//cdn.bootcss.com/jquery/2.2.4/jquery.min.js"></script>
<script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
<script src="//cdn.bootcss.com/angular.js/1.5.7/angular.min.js"></script>
<script src="//cdn.bootcss.com/html2canvas/0.4.1/html2canvas.min.js"></script>
<script type="text/javascript">
var WebSocketIP = {{.ip}}

var MS = new Array()
function test(){
    var messages = new Array()
    var doms = $('.daodream-conversation-part')
    for(i=0;i<doms.length;i++){
        var dom = $($(doms[i]).children()[0])
        var _class = dom.attr("class")
        var who = ""

        var messageDom = $(dom.find('.daodream-comment-body')[0])
        if (_class.indexOf("user") > 0) {
            who = "user"
        }
        if (_class.indexOf("admin") > 0){
            who = "admin"
        }
        var p = messageDom.find('p')
        var message = p.length == 1 ? p.html() : messageDom.html()

        messages.push(who+':::'+message)
    }
    MS = messages
    alert(MS)
}


$(document).ready( function(){
                $(".example1").on("click", function(event) {
                        event.preventDefault();
                        html2canvas(document.getElementById('all'), {
                            allowTaint: true,
                            taintTest: false,
                            onrendered: function(canvas) {
                                canvas.id = "mycanvas";
                                //document.body.appendChild(canvas);
                                //生成base64图片数据
                                var img = new Image()
                                img.crossOrigin = "*";
                                img.onload = function() {
                                    ctx.drawImage(img, 0, 0);
                                    var g = ctx.canvas.toDataURL("image/jpeg", 1);
                                    console.log(g)
                                }
                                var dataUrl = canvas.toDataURL();
                                var newImg = document.createElement("img");
                                newImg.src =  dataUrl;
                                document.body.appendChild(newImg);
                            }
                        });
                });

        });
</script>
<script src="/static/js/index2.js"></script>
</html>