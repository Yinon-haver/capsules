(this["webpackJsonpcapsules-frontend-test"]=this["webpackJsonpcapsules-frontend-test"]||[]).push([[0],{18:function(e,t,n){e.exports=n(19)},19:function(e,t,n){"use strict";n.r(t);var a=n(2),o=n(3),s=n(5),r=n(4),c=n(6),l=n(0),u=n.n(l),p=n(17),i=n.n(p),h=n(7),m=n.n(h),f=(n(41),"https://capsules-web-server.herokuapp.com"),d="ws://capsules-web-server.herokuapp.com",g=function(e){function t(e){var n;return Object(a.a)(this,t),(n=Object(s.a)(this,Object(r.a)(t).call(this,e))).send=function(){m.a.post(f+"/users",{Phone:n.state.Phone,Token:n.state.Token}).then((function(e){console.log(e.request.response)}))},n.state={Phone:"",Token:""},n}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;return u.a.createElement("div",null,u.a.createElement("span",null,"Add user:"),u.a.createElement("input",{onChange:function(t){return e.setState({Phone:t.target.value})},type:"text",placeholder:"phone"}),u.a.createElement("input",{onChange:function(t){return e.setState({Token:t.target.value})},type:"text",placeholder:"token"}),u.a.createElement("button",{onClick:this.send},"send"))}}]),t}(u.a.Component),v=function(e){function t(e){var n;return Object(a.a)(this,t),(n=Object(s.a)(this,Object(r.a)(t).call(this,e))).send=function(){m.a.get(f+"/capsules",{params:{phone:n.state.phone,amount:n.state.amount,offset:n.state.offset,isWatched:n.state.isWatched}}).then((function(e){console.log(e.request.response)}))},n.state={phone:"",amount:0,offset:0,isWatched:!1},n}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;return u.a.createElement("div",null,u.a.createElement("span",null,"Get capsules:"),u.a.createElement("input",{onChange:function(t){return e.setState({phone:t.target.value})},type:"text",placeholder:"phone"}),u.a.createElement("input",{onChange:function(t){return e.setState({amount:+t.target.value})},type:"text",placeholder:"amount"}),u.a.createElement("input",{onChange:function(t){return e.setState({offset:+t.target.value})},type:"text",placeholder:"offset"}),u.a.createElement("input",{onChange:function(t){return e.setState({isWatched:"true"===t.target.value})},type:"text",placeholder:"isWatched"}),u.a.createElement("button",{onClick:this.send},"send"))}}]),t}(u.a.Component),E=function(e){function t(e){var n;return Object(a.a)(this,t),(n=Object(s.a)(this,Object(r.a)(t).call(this,e))).send=function(){m.a.get(f+"/chat",{params:{phone:n.state.phone,capsuleID:n.state.capsuleID,amount:n.state.amount,offset:n.state.offset}}).then((function(e){console.log(e.request.response)}))},n.state={phone:"",capsuleID:0,amount:0,offset:0},n}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;return u.a.createElement("div",null,u.a.createElement("span",null,"Get messages:"),u.a.createElement("input",{onChange:function(t){return e.setState({phone:t.target.value})},type:"text",placeholder:"phone"}),u.a.createElement("input",{onChange:function(t){return e.setState({capsuleID:+t.target.value})},type:"text",placeholder:"capsuleID"}),u.a.createElement("input",{onChange:function(t){return e.setState({amount:+t.target.value})},type:"text",placeholder:"amount"}),u.a.createElement("input",{onChange:function(t){return e.setState({offset:+t.target.value})},type:"text",placeholder:"offset"}),u.a.createElement("button",{onClick:this.send},"send"))}}]),t}(u.a.Component),b=function(e){function t(e){var n;return Object(a.a)(this,t),(n=Object(s.a)(this,Object(r.a)(t).call(this,e))).send=function(){m.a.post(f+"/capsules",{Phone:n.state.Phone,ToPhones:n.state.ToPhones,Content:n.state.Content,OpenDate:n.state.OpenDate}).then((function(e){console.log(e.request.response)}))},n.addParticipate=function(){var e=n.state.ToPhones;e.push("");var t=u.a.createElement("li",{key:n.state.ToPhones.length},u.a.createElement("input",{name:n.state.toPhonesNum,onChange:function(e){var t=n.state.ToPhones;t[e.target.name]=e.target.value,n.setState({ToPhones:t})},type:"text",placeholder:"toPhone"})),a=n.state.phonesTextFields;a.push(t),n.setState({phonesTextFields:a,toPhonesNum:n.state.toPhonesNum+1,ToPhones:e})},n.state={Phone:"",Content:"",OpenDate:"",ToPhones:[""],phonesTextFields:[u.a.createElement("li",{key:0},u.a.createElement("input",{name:"0",onChange:function(e){var t=n.state.ToPhones;t[e.target.name]=e.target.value,n.setState({ToPhones:t})},type:"text",placeholder:"toPhone"}))],toPhonesNum:1},n}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;return u.a.createElement("div",null,u.a.createElement("span",null,"Create capsule:"),u.a.createElement("input",{onChange:function(t){return e.setState({Phone:t.target.value})},type:"text",placeholder:"phone"}),u.a.createElement("input",{onChange:function(t){return e.setState({Content:t.target.value})},type:"text",placeholder:"content"}),u.a.createElement("input",{onChange:function(t){return e.setState({OpenDate:t.target.value})},type:"text",placeholder:"openDate"}),u.a.createElement("button",{onClick:this.addParticipate},"add"),u.a.createElement("ul",null,this.state.phonesTextFields),u.a.createElement("button",{onClick:this.send},"send"))}}]),t}(u.a.Component),C=function(e){function t(e){var n;return Object(a.a)(this,t),(n=Object(s.a)(this,Object(r.a)(t).call(this,e))).connect=function(e,t){n.websocket=new WebSocket(d+"/openChatConnection?phone="+t+"&capsuleID="+e),n.websocket.onmessage=function(e){console.log(e.data)}},n.send=function(){n.state.websocket.send(n.state.message)},n.state={phone:"",capsuleID:"",message:""},n}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){var e=this;return u.a.createElement("div",null,u.a.createElement("span",null,"Open chat connection:"),u.a.createElement("input",{onChange:function(t){return e.setState({capsuleID:t.target.value})},type:"text",placeholder:"capsuleID"}),u.a.createElement("input",{onChange:function(t){return e.setState({phone:t.target.value})},type:"text",placeholder:"phone"}),u.a.createElement("button",{onClick:function(){return e.connect(e.state.capsuleID,e.state.phone)}},"connect"),u.a.createElement("br",null),u.a.createElement("span",null,"Message:"),u.a.createElement("input",{onChange:function(t){return e.setState({message:t.target.value})},type:"text",placeholder:"message"}),u.a.createElement("button",{onClick:this.send},"send"))}}]),t}(u.a.Component);C.websocket=null;var O=function(e){function t(){return Object(a.a)(this,t),Object(s.a)(this,Object(r.a)(t).apply(this,arguments))}return Object(c.a)(t,e),Object(o.a)(t,[{key:"render",value:function(){return u.a.createElement("div",null,u.a.createElement(g,null),u.a.createElement(v,null),u.a.createElement(E,null),u.a.createElement(b,null),u.a.createElement(C,null))}}]),t}(u.a.Component);i.a.render(u.a.createElement(O,null),document.getElementById("root"))},41:function(e,t,n){}},[[18,1,2]]]);
//# sourceMappingURL=main.6eefde2d.chunk.js.map