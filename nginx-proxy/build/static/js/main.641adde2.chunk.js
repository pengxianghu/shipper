(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{11:function(e,t,n){e.exports=n(19)},16:function(e,t,n){},17:function(e,t,n){},19:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n.n(a),o=n(9),s=n.n(o),c=(n(16),n(1)),i=n(2),l=n(4),m=n(3),u=n(5),h=(n(17),n(6)),p=n(10),d=n.n(p),g=function(e){function t(e){var n;return Object(c.a)(this,t),(n=Object(l.a)(this,Object(m.a)(t).call(this,e))).state={created:!1,description:"",weight:0,containers:[],consignments:[]},n.create=function(){var e=n.state,t=sessionStorage.getItem("token");fetch("http://www.pengxianghu.com/rpc",{method:"POST",headers:{"Content-Type":"application/json",Authorization:t},body:JSON.stringify({service:"go.micro.srv.consignment",method:"ConsignmentService.CreateConsignment",request:d.a.omit(e,"created","consignments")})}).then(function(e){return e.json()}).then(function(t){void 0===n.state.consignments?n.setState({created:t.created,consignments:[e]}):n.setState({created:t.created,consignments:[].concat(Object(h.a)(n.state.consignments),[e])})})},n.addContainer=function(e){n.setState({containers:[].concat(Object(h.a)(n.state.containers),[e.target.value])})},n.setDescription=function(e){n.setState({description:e.target.value})},n.setWeight=function(e){n.setState({weight:Number(e.target.value)})},n.state={},n}return Object(u.a)(t,e),Object(i.a)(t,[{key:"componentWillMount",value:function(){var e=this,t=sessionStorage.getItem("token");fetch("http://www.pengxianghu.com/rpc",{method:"POST",headers:{"Content-Type":"application/json",Authorization:t},body:JSON.stringify({service:"go.micro.srv.consignment",method:"ConsignmentService.GetConsignments",request:{}})}).then(function(e){return e.json()}).then(function(t){console.log("create consignment component mount: "+t),e.setState({consignments:t.consignments})})}},{key:"render",value:function(){var e=this.state.consignments;return r.a.createElement("div",{className:"consignment-screen"},r.a.createElement("div",{className:"consignment-form container"},r.a.createElement("br",null),r.a.createElement("div",{className:"form-group"},r.a.createElement("textarea",{onChange:this.setDescription,className:"form-control",placeholder:"Description"})),r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{onChange:this.setWeight,type:"number",placeholder:"Weight",className:"form-control"})),r.a.createElement("br",null),r.a.createElement("button",{onClick:this.create,className:"btn btn-primary"},"\u6dfb\u52a0"),r.a.createElement("br",null),r.a.createElement("hr",null)),!!(e&&e.length>0)&&r.a.createElement("div",{className:"consignment-list"},r.a.createElement("h2",null,"Consignments"),r.a.createElement("hr",null),e.map(function(e){return r.a.createElement("div",null,r.a.createElement("p",null,"Description: ",e.description),r.a.createElement("p",null,"Weight: ",e.weight),r.a.createElement("hr",null))})))}}]),t}(r.a.Component),f=function(e){function t(e){var n;return Object(c.a)(this,t),(n=Object(l.a)(this,Object(m.a)(t).call(this,e))).state={authenticated:!1,email:"",password:"",err:""},n.login=function(){fetch("http://www.pengxianghu.com/rpc",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({service:"go.micro.srv.user",method:"UserService.Auth",request:{email:n.state.email,password:n.state.password}})}).then(function(e){return e.json()}).then(function(e){void 0!==e.token?(n.setState({token:e.token,authenticated:!0}),sessionStorage.setItem("token",e.token),n.props.onAuth(e.token)):console.log("auth failed!")}).catch(function(e){return n.setState({err:e,authenticated:!1})})},n.signup=function(){fetch("http://www.pengxianghu.com/rpc",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({service:"go.micro.srv.user",method:"UserService.Create",request:{email:n.state.email,password:n.state.password,name:n.state.name}})}).then(function(e){return e.json()}).then(function(e){}).catch(function(e){return n.setState({err:e,authenticated:!1})})},n.setEmail=function(e){n.setState({email:e.target.value})},n.setPassword=function(e){n.setState({password:e.target.value})},n.setName=function(e){n.setState({name:e.target.value})},n.state={},n}return Object(u.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){return r.a.createElement("div",{className:"Authenticate"},r.a.createElement("div",{className:"Login"},r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{type:"email",onChange:this.setEmail,placeholder:"E-Mail",className:"form-control"})),r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{type:"password",onChange:this.setPassword,placeholder:"Password",className:"form-control"})),r.a.createElement("button",{className:"btn btn-primary",onClick:this.login},"\u767b\u5f55"),r.a.createElement("br",null),r.a.createElement("br",null)),r.a.createElement("div",{className:"Sign-up"},r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{type:"input",onChange:this.setName,placeholder:"Name",className:"form-control"})),r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{type:"email",onChange:this.setEmail,placeholder:"E-Mail",className:"form-control"})),r.a.createElement("div",{className:"form-group"},r.a.createElement("input",{type:"password",onChange:this.setPassword,placeholder:"Password",className:"form-control"})),r.a.createElement("button",{className:"btn btn-primary",onClick:this.signup},"\u6ce8\u518c")))}}]),t}(r.a.Component),v=function(e){function t(){var e,n;Object(c.a)(this,t);for(var a=arguments.length,o=new Array(a),s=0;s<a;s++)o[s]=arguments[s];return(n=Object(l.a)(this,(e=Object(m.a)(t)).call.apply(e,[this].concat(o)))).state={err:null,authenticated:!1},n.onAuth=function(e){n.setState({authenticated:!0})},n.renderLogin=function(){return r.a.createElement(f,{onAuth:n.onAuth})},n.renderAuthenticated=function(){return r.a.createElement(g,null)},n.getToken=function(){return sessionStorage.getItem("token")||!1},n.isAuthenticated=function(){return n.state.authenticated||n.getToken()||!1},n}return Object(u.a)(t,e),Object(i.a)(t,[{key:"render",value:function(){var e=this.isAuthenticated();return r.a.createElement("div",{className:"App"},r.a.createElement("div",{className:"App-header"},r.a.createElement("h2",null,"Shipper Golang Micro-Service")),r.a.createElement("div",{className:"App-intro container"},e?this.renderAuthenticated():this.renderLogin()))}}]),t}(a.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));s.a.render(r.a.createElement(v,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}},[[11,1,2]]]);
//# sourceMappingURL=main.641adde2.chunk.js.map