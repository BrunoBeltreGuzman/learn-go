_N_E=(window.webpackJsonp_N_E=window.webpackJsonp_N_E||[]).push([[8],{"HaE+":function(e,t,n){"use strict";function c(e,t,n,c,r,s,o){try{var i=e[s](o),a=i.value}catch(u){return void n(u)}i.done?t(a):Promise.resolve(a).then(c,r)}function r(e){return function(){var t=this,n=arguments;return new Promise((function(r,s){var o=e.apply(t,n);function i(e){c(o,r,s,i,a,"next",e)}function a(e){c(o,r,s,i,a,"throw",e)}i(void 0)}))}}n.d(t,"a",(function(){return r}))},s75Z:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/getFetch",function(){return n("wM0U")}])},wM0U:function(e,t,n){"use strict";n.r(t),n.d(t,"default",(function(){return u}));var c=n("nKUr"),r=n("o0o1"),s=n.n(r),o=n("HaE+"),i=n("q1tI"),a=n("7oih");function u(){var e=Object(i.useState)([]),t=e[0],n=e[1],r=Object(i.useState)(!0),u=r[0],d=r[1],j=Object(i.useState)([]),l=j[0],h=j[1];return Object(i.useEffect)(Object(o.a)(s.a.mark((function e(){var t,c;return s.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,d(!0),e.next=4,fetch("https://jsonplaceholder.typicode.com/posts");case 4:return t=e.sent,e.next=7,t.json();case 7:c=e.sent,console.log(c),n(c),d(!1),e.next=18;break;case 13:e.prev=13,e.t0=e.catch(0),console.log(e.t0),d(!1),h(e.t0);case 18:case"end":return e.stop()}}),e,null,[[0,13]])}))),[]),u?Object(c.jsx)(a.a,{children:Object(c.jsx)("div",{className:"row p-5",children:Object(c.jsx)("div",{className:"col-md-10 m-auto text-center",children:Object(c.jsx)("h1",{children:"Loading..."})})})}):t?Object(c.jsx)(a.a,{children:Object(c.jsx)("div",{className:"row p-5",children:Object(c.jsxs)("div",{className:"col-md-10 m-auto text-center",children:[Object(c.jsx)("h1",{children:"Data!"}),t.map((function(e,t){return Object(c.jsx)("pre",{children:Object(c.jsx)("p",{children:JSON.stringify(e)})},t)}))]})})}):l?Object(c.jsx)(a.a,{children:Object(c.jsx)("div",{className:"row p-5",children:Object(c.jsxs)("div",{className:"col-md-10 m-auto text-center",children:[Object(c.jsx)("h1",{children:"Error"}),Object(c.jsx)("pre",{children:JSON.stringify(l)})]})})}):void 0}}},[["s75Z",0,2,1,3]]]);