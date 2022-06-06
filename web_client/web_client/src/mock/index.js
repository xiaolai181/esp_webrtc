import Mock from 'mockjs'

const testData=Mock.mock('http://localhost:8080/test','get',{
	status:200, //请求成功状态码
	dataList:[1,2,3,4,5,6,7,8,9,10] //模拟的请求数据
})

//导出
export default testData