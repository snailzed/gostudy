(function(t){function e(e){for(var n,s,l=e[0],i=e[1],c=e[2],p=0,m=[];p<l.length;p++)s=l[p],a[s]&&m.push(a[s][0]),a[s]=0;for(n in i)Object.prototype.hasOwnProperty.call(i,n)&&(t[n]=i[n]);u&&u(e);while(m.length)m.shift()();return o.push.apply(o,c||[]),r()}function r(){for(var t,e=0;e<o.length;e++){for(var r=o[e],n=!0,l=1;l<r.length;l++){var i=r[l];0!==a[i]&&(n=!1)}n&&(o.splice(e--,1),t=s(s.s=r[0]))}return t}var n={},a={app:0},o=[];function s(e){if(n[e])return n[e].exports;var r=n[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,s),r.l=!0,r.exports}s.m=t,s.c=n,s.d=function(t,e,r){s.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},s.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},s.t=function(t,e){if(1&e&&(t=s(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)s.d(r,n,function(e){return t[e]}.bind(null,n));return r},s.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return s.d(e,"a",e),e},s.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},s.p="/";var l=window["webpackJsonp"]=window["webpackJsonp"]||[],i=l.push.bind(l);l.push=e,l=l.slice();for(var c=0;c<l.length;c++)e(l[c]);var u=i;o.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("56d7")},1015:function(t,e,r){},"10dd":function(t,e,r){"use strict";var n=r("188e"),a=r.n(n);a.a},"188e":function(t,e,r){},"2f34":function(t,e,r){},"56d7":function(t,e,r){"use strict";r.r(e);r("cadf"),r("551c"),r("097d");var n=r("2b0e"),a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{attrs:{id:"app"}},[r("router-view")],1)},o=[],s={name:"App"},l=s,i=(r("8840"),r("2877")),c=Object(i["a"])(l,a,o,!1,null,null,null);c.options.__file="App.vue";var u=c.exports,p=r("8c4f"),m=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"layout-page"},[r("div",{staticClass:"main-page"},[r("router-view")],1)])},f=[],d={name:"Layout"},b=d,v=Object(i["a"])(b,m,f,!1,null,null,null);v.options.__file="index.vue";var _=v.exports,h=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"home"},[r("Tabs",{attrs:{value:"1"}},[r("TabPane",{attrs:{label:"技术",name:"1"}},t._l(t.list1,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])})),r("TabPane",{attrs:{label:"生活",name:"2"}},t._l(t.list2,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])})),r("TabPane",{attrs:{label:"八卦",name:"3"}},t._l(t.list3,function(e){return r("Card",{key:e.id,attrs:{title:e.title}},[r("p",[t._v(t._s(e.content))]),r("p",[t._v("提问时间："+t._s(e.createAt)+" 提问者："+t._s(e.createByName))])])}))],1)],1)},y=[],x={name:"home",data:function(){return{list1:[{id:1,title:"aaaa",content:"bbbbbbbbbb",createAt:"2018-10-10",createByName:"rrrrr"}],list2:[{id:1,title:"aaaa",content:"bbbbbbbbbb",craeteAt:"2018-10-10",craeteByName:"rrrrr"}],list3:[{id:1,title:"aaaa",content:"bbbbbbbbbb",craeteAt:"2018-10-10",craeteByName:"rrrrr"}]}}},g=x,k=Object(i["a"])(g,h,y,!1,null,null,null);k.options.__file="index.vue";var w=k.exports,I=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"登录"}},[r("Form",[r("FormItem",{attrs:{prop:"username"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("登录")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/register"}},[t._v("注册")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},$=[],O={name:"Login",data:function(){return{form:{user:"",password:""}}},methods:{submit:function(){console.log("调用登录接口")}}},C=O,j=(r("913c"),Object(i["a"])(C,I,$,!1,null,"50f388e9",null));j.options.__file="index.vue";var F=j.exports,P=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"login"},[r("div",{staticClass:"login-box"},[r("Card",{attrs:{title:"注册"}},[r("Form",[r("FormItem",{attrs:{prop:"sex"}},[r("RadioGroup",{model:{value:t.form.sex,callback:function(e){t.$set(t.form,"sex",e)},expression:"form.sex"}},[r("Radio",{attrs:{label:1}},[t._v("男")]),r("Radio",{attrs:{label:2}},[t._v("女")])],1)],1),r("FormItem",{attrs:{prop:"nickname"}},[r("Input",{attrs:{type:"text",placeholder:"昵称"},model:{value:t.form.nickname,callback:function(e){t.$set(t.form,"nickname",e)},expression:"form.nickname"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-chatbubbles-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"email"}},[r("Input",{attrs:{type:"text",placeholder:"Email"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-mail-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"user"}},[r("Input",{attrs:{type:"text",placeholder:"Username"},model:{value:t.form.user,callback:function(e){t.$set(t.form,"user",e)},expression:"form.user"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1),r("FormItem",{attrs:{prop:"password"}},[r("Input",{attrs:{type:"password",placeholder:"Password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}},[r("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1),r("Button",{attrs:{type:"primary",long:""},on:{click:t.submit}},[t._v("注册")]),r("router-link",{staticClass:"ivu-btn ivu-btn-default ivu-btn-long",staticStyle:{"margin-top":"10px"},attrs:{to:"/login"}},[t._v("登录")]),r("div",{staticStyle:{"margin-top":"10px","text-align":"right"}},[r("router-link",{staticClass:"ivu-btn ivu-btn-text",attrs:{to:"/"}},[t._v("返回首页")])],1)],1)],1)],1)])},S=[],A=(r("96cf"),r("1da1")),B={name:"Regiter",data:function(){return{form:{user:"",nickname:"",sex:1,email:"",password:""}}},methods:{submit:function(){var t=Object(A["a"])(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.$http.post("/api/user/register",this.form);case 2:e=t.sent,0===e.code&&(this.$Message.success("注册成功"),this.$router.push("/login"));case 4:case"end":return t.stop()}},t,this)}));return function(){return t.apply(this,arguments)}}()}},T=B,E=(r("10dd"),Object(i["a"])(T,P,S,!1,null,"c51b04c4",null));E.options.__file="index.vue";var R=E.exports,N=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("Card",{attrs:{title:"提问"}},[r("Form",{attrs:{"label-position":"top",model:t.form}},[r("FormItem",{attrs:{prop:"title",label:"标题"}},[r("Input",{attrs:{type:"text",placeholder:"请输入标题"},model:{value:t.form.title,callback:function(e){t.$set(t.form,"title",e)},expression:"form.title"}})],1),r("FormItem",{attrs:{prop:"category",label:"类别"}},[r("Select",{attrs:{placeholder:"请选择类别"},model:{value:t.form.category,callback:function(e){t.$set(t.form,"category",e)},expression:"form.category"}},[r("Option",{attrs:{value:"1"}},[t._v("技术")]),r("Option",{attrs:{value:"2"}},[t._v("生活")]),r("Option",{attrs:{value:"3"}},[t._v("搞笑")])],1)],1),r("FormItem",{attrs:{prop:"centent",label:"内容"}},[r("Input",{attrs:{type:"textarea",placeholder:"请输入内容"},model:{value:t.form.centent,callback:function(e){t.$set(t.form,"centent",e)},expression:"form.centent"}})],1),r("Button",{attrs:{type:"primary"},on:{click:t.submit}},[t._v("提交")])],1)],1)],1)},M=[],U={name:"Ask",data:function(){return{form:{title:""}}}},L=U,J=(r("8d2a"),Object(i["a"])(L,N,M,!1,null,"dd3a215c",null));J.options.__file="index.vue";var G=J.exports;n["default"].use(p["a"]);var q=new p["a"]({routes:[{path:"/",component:_,children:[{path:"",name:"home",component:w},{path:"ask",name:"ask",component:G}]},{path:"/login",name:"login",component:F},{path:"/register",name:"register",component:R}]}),z=r("2f62");n["default"].use(z["a"]);var D=new z["a"].Store({state:{},mutations:{},actions:{}}),H=r("e069"),K=r.n(H),Q=(r("dcad"),r("bc3a")),V=r.n(Q),W=V.a.create({baseURL:"localhost:9090",timeout:1e4,headers:{"Content-Type":"application/json;charset=UTF-8"}}),X=W;n["default"].use(K.a),n["default"].prototype.$http=X,n["default"].config.productionTip=!1,new n["default"]({el:"#app",router:q,store:D,render:function(t){return t(u)}})},8840:function(t,e,r){"use strict";var n=r("b714"),a=r.n(n);a.a},"8d2a":function(t,e,r){"use strict";var n=r("1015"),a=r.n(n);a.a},"913c":function(t,e,r){"use strict";var n=r("2f34"),a=r.n(n);a.a},b714:function(t,e,r){}});
//# sourceMappingURL=app.f280d791.js.map