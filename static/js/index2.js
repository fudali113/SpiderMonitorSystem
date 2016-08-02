var monitor = angular.module('monitor', []);

monitor.directive('spider', function() {
    return {
        restrict: 'E',
        templateUrl: '/static/html/spider.html',
        replace: true
    };
});

monitor.directive('computer', function() {
    return {
        restrict: 'E',
        templateUrl: '/static/html/computer05.html',
        replace: true,
        link:function(scope, el, attr) { 
		
		} 
    };
});

monitor.directive('smss',function(){
	return {
		restrict: 'E',
		replace: true,
		link:function(scope){
			
		},
		templateUrl: '/static/html/smss.html'
	}
})

monitor.directive('setting',function(){
	return {
		restrict: 'E',
		replace: true,
		link:function(scope){
			$('.alert').hide()
			scope.setting = function(){
				
			}
		},
		templateUrl: '/static/html/setting.html'
	}
})

monitor.service( 'computer', [ '$rootScope', function( $rootScope ) {
    var service = {
      	computers: [],
      	addComputer: function (data) {
      		
  			var nowComputer = undefined
        		for (var i = this.computers.length - 1; i >= 0; i--) {
				if(data.pc_id == this.computers[i].id){
					nowComputer = i 
				}
			}
			if (nowComputer != undefined) {
				if (data.sys != undefined){
					this.computers[nowComputer].sys.cpu = data.sys.cpu[0]
					this.computers[nowComputer].sys.mem = data.sys.mem.usedpercent
					this.computers[nowComputer].sys.proc = data.sys.proc
					return
				}
				if (data.hb == -1){
					this.computers.splice(newComputer,1)
				}else{
					if (data.bank_status != undefined) {
						this.computers[nowComputer].receiveData(data)
						this.computers[nowComputer].hb = 1
					}
					if (data.hb != undefined) {
						this.computers[nowComputer].hb = data.hb
					}
					if (data.ip != undefined) {
						this.computers[nowComputer].ip = data.ip
					}
					this.setHbStyle(nowComputer)
				}
			}else{
					var one = newComputer(data.pc_id)
      				if (data.bank_status != undefined) {
						one.receiveData(data)
						one.hb = 1
      				}
					this.setHbStyle(this.computers.push(one)-1)
			}
        	$rootScope.$broadcast( 'computers.update' );
      	},
      	setHbStyle:function(i){
      		var hb = this.computers[i].hb
      		if (hb == 0) {
				this.computers[i].hbStyle.background = "red"
			}else if (hb == 1) {
				this.computers[i].hbStyle.background = "#00FF7F"
			}
      	}
    }
    var websocket = new WebSocket("ws://"+WebSocketIP+"/ws");
	websocket.onopen = function(evt) { 
        alert('websocket连接成功') 
    }; 
    websocket.onclose = function(evt) { 
        alert('websocket断开连接') 
    }; 
    websocket.onmessage = function(evt) {
		var data = JSON.parse(evt.data)
		if(data instanceof Array ){
			for (var i=0;i<data.length;i++){
				service.addComputer(data[i])
			}
		}else{
			service.addComputer(data)
		}
    }; 
    websocket.onerror = function(evt) { 
        alert('websocket出现错误') 
    }; 
    return service
}]);

monitor.filter('hb_filter',function(){
	return function(input){
		var r = input == 0 ? 0 : 1 
		return r
	}
})

monitor.controller('computers',['$scope','$http','computer',function($scope,$http,computer){

	$scope.$on( 'computers.update', function( event ) {
        $scope.computers = computer.computers;
		var activeNum = 0
		var total = computer.computers.length
		for (var i = total - 1; i >= 0; i--) {
			var c = computer.computers[i]
			if (c.hb == 0) {
				$('#'+c.id).collapse('hide')
			}else{
				activeNum++
				$('#'+c.id).collapse('show')
			}
		}
		$scope.active = activeNum
		$scope.total = total
        $scope.$apply();
    }); 
	$scope.computers = computer.computers 
	$scope.showDetails = false
	$scope.total = 0
	$scope.active = 0
	$scope.settingRequestAnimation = false
    $scope.param = {}
	
	$http({
		url:'/setting',
		method:'get',
	}).success(function(data){
		$scope.param = data
		theme = data.theme
	}).error(function(data){
		alert("error")
	});
	

    $scope.submitSetting = function(){
		$scope.settingRequestAnimation = true
    		$http({
			url:'/setting',
			method:'post',
			data:$scope.param
		}).success(function(data){
			$scope.param = data
			setTimeout("",500)
			if(data.theme != theme){
					window.location.reload();
			}
			$scope.settingRequestAnimation = false
		}).error(function(data){
			alert("error")
		});
    }
    $scope.submitDefalutSetting = function(){
		$scope.settingRequestAnimation = true
    		$http({
			url:'/setting/default',
			method:'post'
		}).success(function(data){
			if(data){
				$scope.param = data
				setTimeout("",1000)
				if(data.theme != theme){
					window.location.reload();
				}
				$scope.settingRequestAnimation = false
			}
		}).error(function(data){
			alert("error")
		});
    }
	$scope.HideSettingAnimation = function(data){
		if(data.theme != theme){
			window.location.reload();
		}
		$scope.settingRequestAnimation = false
	}
	$scope.hideAttr = function(k){
	    	var hides = ['sid','bid','step','bank_name','stc']
	    	for (var i = hides.length - 1; i >= 0; i--) {
	    		if (k == hides[i]) {
	    			return true
	    		}
	    	}
	    	return false
	    }

	
}])

var newComputer = function(cid){
	var computer = {
		hb:1, //heartbeat
		hbStyle:{
			background:"blue"
		},
		sys:{
			cpu:0,
			mem:0,
			proc:0
		},
		ip:undefined,
		id:cid,
		spiders:[],
		receiveData : function(data){
			this.ip = data.ip
			this.id = data.pc_id
			for(var i=0;i<this.spiders.length;i++){
				if(this.spiders[i].bank_name == data.bank_name){
					this.spiders[i].content = data.bank_status
					return
				}
			}
			this.spiders.push(newSpider(data.bank_name,data.bank_status))
		}
	}
	return computer
}

var newSpider = function(bid,data){
	return {
		bank_name:bid,
		content:data
	}
}
