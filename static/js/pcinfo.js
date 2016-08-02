var system = angular.module('system', []);

system.directive('cpu', function() {
    return {
		scope:{
			id:"@",
			date:"=",
			data:"="
			},
        restrict: 'E',
        template: '<div style="height:300px;width:400px"></div>',
        replace: true,
		link: function($scope, element, attrs) {  
			var getOption = function(){
				return {
				    xAxis: {
				        type: 'category',
				        boundaryGap: false,
				        data: $scope.date
				    },
				    yAxis: {
				        boundaryGap: [0, '50%'],
				        type: 'value'
				    },
				    series: [
				        {
				            name:'成交',
				            type:'line',
				            smooth:true,
				            symbol: 'none',
				            stack: 'a',
				            areaStyle: {
				                normal: {}
				            },
				            data: $scope.data
				        }
				    ]
				};
			}
			var dom = echarts.init(document.getElementById($scope.id)); 
			dom.setOption(getOption())
			$scope.$on( 'sysinfo.update.link', function( event ) {
				dom.setOption(getOption())
		    }); 
		}
    };
});

system.service( 'sysinfo', [ '$rootScope','$http', function( $rootScope,$http ) {
	var service = {
		data:{
			date:[],
			data:[]
		},
		addData:function(data){
			var d = new Date()
			var getDouble = function(i){
				return i < 10 ? '0'+i : i
			}
			this.data.date.push(getDouble(d.getHours())+':'+getDouble(d.getMinutes())+':'+getDouble(d.getSeconds()) )
			this.data.data.push(data.cpu[0])
			$rootScope.$broadcast( 'sysinfo.update' );
		}
	}
	
	var getSysinfo = function(){
		$http({
			url:'/'+pcid+'/info/all',
			method:'get',
		}).success(function(data){
			service.addData(data)
		})
	}
	setInterval(getSysinfo,5000)
	return service
}]);

system.controller('sysinfoshow',['$rootScope','$scope','$http','sysinfo',function($rootScope,$scope,$http,sysinfo){
	$scope.$on( 'sysinfo.update', function( event ) {
        $scope.data = sysinfo.data.data
		$scope.date = sysinfo.data.date
		$rootScope.$broadcast( 'sysinfo.update.link' );
    }); 
	$scope.data=[]
	$scope.date=[]
}]);