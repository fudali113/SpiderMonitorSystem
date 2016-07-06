var monitor = angular.module('monitor', []);
monitor.directive('computer', function() {
    return {
        restrict: 'E',
        templateUrl: '/static/computer.html',
        replace: true,
        link:function(scope, el, attr) {  
        	$(".miaoshu div").css("display","none");

        	for (var i = 6; i > 0; i--) {
        		$(".bfb"+i+" img").animate({width:+"0px"})	
        	}
        	
            scope.showDetails =   function(event,i){
				$(event.target).parent().parent().find(".miaoshu .miaoshu"+i).toggle(500).siblings().fadeOut();
			}
        }
    };
});

monitor.service( 'computer', [ '$rootScope', function( $rootScope ) {
    var service = {
      	computers: [],
      	addComputer: function (data) {
      		if (this.computers.length == 0) {
      				var one = newComputer(data.pc_id)
					one.receiveData(data.bank_status)
					this.computers.push(one)
      		}else{
      			var nowComputer = undefined
	        	for (var i = this.computers.length - 1; i >= 0; i--) {
					if(data.pc_id == this.computers[i].cid){
						nowComputer = i 
					}
				}
				if (nowComputer != undefined) {
					this.computers[nowComputer].receiveData(data.bank_status)
				}else{
						var one = newComputer(data.pc_id)
						one.receiveData(data.bank_status)
						this.computers.push(one)
				}
			}
        	$rootScope.$broadcast( 'computers.update' );
      	}
    }
    var websocket = new WebSocket("ws://{{.ip}}:8080/ws");
	websocket.onopen = function(evt) { 
            alert('websocket连接成功') 
        }; 
    websocket.onclose = function(evt) { 
            alert('websocket断开连接') 
        }; 
    websocket.onmessage = function(evt) {
			service.addComputer(JSON.parse(evt.data))
        }; 
    websocket.onerror = function(evt) { 
            alert('websocket出现错误') 
        }; 
    return service
}]);

monitor.controller('computers',['$scope','computer',function($scope,computer){

	$scope.$on( 'computers.update', function( event ) {
        $scope.computers = computer.computers;
        $scope.$apply();
    }); 

	$scope.computers = computer.computers  //[newComputer("1234")]

	$scope.showDetails = function(){
		$(this).parent().parent().find(".miaoshu .miaoshu2").toggle(500).siblings().fadeOut();
	}

	
}])

var newComputer = function(cid){
	var computer = {
		s:0,
		hb:1, //heartbeat
		cid:cid,
		sid:undefined,
		bid:undefined,
		hository:new Array(6),
		receiveData : function(data){
			if (typeof data == 'string'){
				var data = JSON.parse(data)
			}
			if(data.step == 0 || this.sid != data.sid){
					this.sid = data.sid
					this.bid = data.bid
					$('#sid').html(data.sid)
					$('#bid').html(data.bid)
					this.hository = new Array(6)
					this.initFlowSheet()
				}
			this.s = data.step
			this.hository[data.step] = data
			this.loadBegin()
			this.startStep(88)
			this.loadMessage()
		},
		startStep : function(bfb){
			var i = this.s
			$(".miaoshu"+i+" .span"+i).text(bfb);
			w =bfb*140/100;
			$(".bfb"+ i +" img").animate({width:w+"px"},bfb*100)
		},
		loadBegin : function(){
			for(i=0;i<this.s;i++){
				$(".bfb"+i+" img").stop(true,true);
				$(".bfb"+i+" img").animate({width:140+"px"},"fast")
				$(".miaoshu"+i+" .span"+i).text(100);
			}
		},
		loadMessage : function(){

			for(var i = 0 ; i <= 6 ; i++){
				var data = this.hository[i]
				var html = '<font>%</font>'
				if(data == undefined){

				}else{
					for (var j in data){
						if(j == 'step'){
							
						}else{
						html += '<br><span contenteditable="true" ><font color="#030303">'+j+'</font> : <font color="#EE4000">'+data[j]+'</font></span>'
					}
				}
			}
				$(".miaoshu"+i+" .span"+i).nextAll().remove();
				$(".miaoshu"+i+" .span"+i).after(html);
			}
		},
		initFlowSheet : function(){
		
			for(var i=1 ; i<=6 ; i++){
				$(".miaoshu"+i+" .span"+i).text(0);
			}
			
			var bfb = $(".miaoshu1 .span1").text();
			 bfb =bfb*140/100;
			$(".bfb1 img").animate({width:bfb+"px"})
			
		 	var bfb = $(".miaoshu2 .span2").text();
			 bfb =bfb*140/100;
			$(".bfb2 img").animate({width:bfb+"px"})
			
			var bfb = $(".miaoshu3 .span3").text();
			 bfb =bfb*140/100;
			$(".bfb3 img").animate({width:bfb+"px"})
			
			var bfb = $(".miaoshu4 .span4").text();
			 bfb =bfb*140/100;
			$(".bfb4 img").animate({width:bfb+"px"})
			
			var bfb = $(".miaoshu5 .span5").text();
			 bfb =bfb*140/100;
			$(".bfb5 img").animate({width:bfb+"px"})
			
			var bfb = $(".miaoshu6 .span6").text();
			 bfb =bfb*140/100;
			$(".bfb6 img").animate({width:bfb+"px"})
		}
	}
	return computer
}