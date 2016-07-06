var monitor = angular.module('monitor', []);
monitor.directive('computer', function() {
    return {
        restrict: 'E',
        templateUrl: '/static/html/computer.html',
        replace: true,
        link:function(scope, el, attr) {  
        	$(".miaoshu div").css("display","none");
            scope.showDetails =   function(event,i){
				$(event.target).parent().parent().parent().find(".miaoshu .miaoshu"+i).toggle(500).siblings().fadeOut();
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
    var websocket = new WebSocket("ws://"+WebSocketIP+":8080/ws");
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
		nowData:{},
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
			if (data.step < this.s) {}else{this.startStep(88)}
		},
		startStep : function(bfb){
			var i = this.s
			$("#"+this.cid+"_percent_"+i).text(bfb);
			w =bfb*140/100;
			$("#"+this.cid+"_animate_img_"+i).animate({width:w+"px"},bfb*100)
		},
		loadBegin : function(){
			for(i=0;i<this.s;i++){
				$("#"+this.cid+"_animate_img_"+i).stop(true,true);
				$("#"+this.cid+"_animate_img_"+i).animate({width:140+"px"},"fast")
				$("#"+this.cid+"_percent_"+i).text(100);
			}
		},
		initFlowSheet : function(){
		
			for(var i=1 ; i<=6 ; i++){
				$("#"+this.cid+"_animate_img_"+i).animate({width:0+"px"})	
				$("#"+this.cid+"_percent_"+i).text(0);
			}
		}
	}
	return computer
}