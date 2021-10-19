(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-871750f0"],{"3d4d":function(e,t,r){"use strict";r.r(t);var o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",[r("Form",{ref:"roleForm",attrs:{model:e.role,rules:e.ruleValidate,"label-position":"right","label-width":100}},[r("Row",{attrs:{gutter:2}},[r("h4",[e._v("角色信息")]),r("Col",{attrs:{span:"7",push:"1"}},[r("FormItem",{attrs:{label:"租户："}},[r("Select",{staticStyle:{width:"170px"},attrs:{prefix:"md-transgender"},model:{value:e.role.v2,callback:function(t){e.$set(e.role,"v2",t)},expression:"role.v2"}},e._l(e.domainList,(function(t){return r("Option",{key:t.value,attrs:{value:t.value}},[e._v(e._s(t.label))])})),1)],1)],1),r("Col",{attrs:{span:"6"}},[r("FormItem",{attrs:{prop:"v5",label:"角色名称："}},[r("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入角色名称",clearable:""},model:{value:e.role.v5,callback:function(t){e.$set(e.role,"v5",t)},expression:"role.v5"}})],1)],1),r("Col",{attrs:{span:"6"}},[r("FormItem",{attrs:{prop:"v1",label:"角色key："}},[r("Input",{staticStyle:{width:"auto"},attrs:{type:"text",prefix:"ios-phone-portrait",placeholder:"请输入角色key",clearable:""},model:{value:e.role.v1,callback:function(t){e.$set(e.role,"v1",t)},expression:"role.v1"}})],1)],1)],1),r("Row",[r("Col",{attrs:{span:"6",push:"4"}},[r("Card",{attrs:{title:"关联菜单",bordered:!1}},[r("Scroll",[r("Tree",{ref:"menuTree",attrs:{data:e.menuList,"show-checkbox":""}})],1)],1)],1),r("Col",{attrs:{span:"6",push:"6"}},[r("Card",{style:{background:"rgba(200,200,200,.2)"},attrs:{"dis-hover":"",bordered:!1}},[r("Tooltip",{attrs:{slot:"title","max-width":"350"},slot:"title"},[r("p",[r("Icon",{attrs:{type:"ios-information-circle-outline",size:"18"}}),e._v("关联权限")],1),r("div",{attrs:{slot:"content"},slot:"content"},[r("h3",[e._v("请在权限管理中设置角色的权限")]),r("h6",[e._v("由于casbin原生的权限模型设计，这里不增加或删除权限")])])]),r("Scroll",[r("Tree",{ref:"policyTree",attrs:{data:e.policyList,"show-checkbox":""}})],1)],1)],1)],1),r("h4",[e._v("其他信息")]),r("br"),r("FormItem",{style:{padding:"0 400px"}},[r("Row",[r("Col",{attrs:{span:"8"}},[r("Button",{attrs:{type:"success",icon:"md-checkmark"},on:{click:e.handleSubmit}},[e._v("保存")])],1),r("Col",{attrs:{span:"8"}},[r("Button",{attrs:{type:"warning",icon:"md-close"},on:{click:e.handleCloseTag}},[e._v("关闭")])],1)],1)],1)],1)],1)},n=[],l=(r("8e6e"),r("456d"),r("ac6a"),r("c5f6"),r("bd86")),i=r("2f62"),a=r("f825");function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,o)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){Object(l["a"])(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}var u={props:{tRole:{type:Object,default:function(){return null}},routeName:String,isModify:Boolean},components:{Tooltip:a["Tooltip"]},name:"OneRolePage",data:function(){return{actionList:[{label:"GET",value:"GET"},{label:"POST",value:"POST"},{label:"PUT",value:"PUT"},{label:"DELETE",value:"DELETE"}],domainList:[],role:{pType:"g"},ruleValidate:{v1:[{required:!0,message:"角色key不能为空",trigger:"blur"}],v5:[{required:!0,message:"角色名称不能为空",trigger:"blur"}]},rom:[],menuList:[],rop:[],policyList:[]}},methods:c({},Object(i["d"])(["closeTag"]),{},Object(i["b"])(["handleDomainList","handleAllMenu","handleRoleOfMenus","handlePolicyAll","handleRoleOfPolicys"]),{handleSubmit:function(){var e=this;this.$refs.roleForm.validate((function(t){if(t){var r={role:c({},e.role)};r.role.id&&(r.role.id=Number(r.role.id));var o=[],n=[],l=e.$refs.menuTree.getCheckedNodes();e.isModify||e.tRole?0===l.length?r.deleteMids=e.tRole.menuList.map((function(e){return e.mid})):(l.forEach((function(t){var r=e.tRole.menuList.some((function(e){return t.id===e.mid}));r||o.push(t.id)})),e.tRole.menuList.forEach((function(e){var t=l.some((function(t){return e.mid===t.id}));t||n.push(e.mid)})),r.addMids=o,r.deleteMids=n):(l.forEach((function(e){return o.push(e.id)})),r.addMids=o),e.$emit("submit",r)}}))},handleReset:function(){this.role={pType:"g"}},handleCloseTag:function(){var e={name:this.routeName};this.tRole&&(e.query=this.tRole),this.closeTag(e)},pullDomainList:function(){var e=this;this.handleDomainList().then((function(t){t&&t.forEach((function(t){return e.domainList.push({value:t,label:t})}))})).catch((function(e){return console.log("错误：",e)}))},pullAllMenu:function(){var e=this;this.handleAllMenu().then((function(t){if(t){var r=[];e.isModify&&e.tRole&&(r=e.tRole.menuList||[]),t.forEach((function(t){var o=[];t.children&&t.children.forEach((function(e){var t=!1;r.length>0&&(t=r.some((function(t){return t.mid===e.id}))),o.push({id:e.id,checked:t,title:e.meta.title||""})}));var n=r.some((function(e){return e.mid===t.id}));e.menuList.push({id:t.id,checked:n,title:t.meta.title||"",expand:!0,children:o})}))}})).catch((function(e){return console.log("错误：",e)}))},pullAllPolicy:function(){var e=this;this.handlePolicyAll().then((function(t){t&&t.forEach((function(t){var r=!1;if(e.isModify&&e.tRole){var o=e.tRole.policyList;r=o.some((function(e){return e.id===t.id}))}e.policyList.push({id:t.id,checked:r,title:t.v5||"",disabled:!0})}))})).catch((function(e){return console.log("错误：",e)}))}}),created:function(){this.tRole&&(this.role=c({},this.tRole)),this.pullDomainList(),this.pullAllMenu(),this.pullAllPolicy()}},d=u,p=(r("4ce6"),r("2877")),f=Object(p["a"])(d,o,n,!1,null,"2c2f5841",null);t["default"]=f.exports},"4ce6":function(e,t,r){"use strict";var o=r("539a"),n=r.n(o);n.a},"539a":function(e,t,r){},e420:function(e,t,r){"use strict";r.r(t);var o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("OneRolePage",{ref:"orp",attrs:{"route-name":"modify_role_page","t-role":e.routeQuery,"is-modify":!0},on:{submit:e.save}})},n=[],l=(r("8e6e"),r("ac6a"),r("456d"),r("bd86")),i=r("3d4d"),a=r("2f62");function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,o)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){Object(l["a"])(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}var u={components:{OneRolePage:i["default"]},name:"add_role_page",data:function(){return{routeQuery:null}},methods:c({},Object(a["b"])(["handleUpdateRole"]),{save:function(e){var t=this;this.handleUpdateRole(e).then((function(e){t.$refs.orp.handleCloseTag(),t.$Message.success("更新角色信息成功")})).catch((function(){return t.$Message.error("更新角色信息失败")}))}}),created:function(){this.routeQuery=this.$route.query}},d=u,p=r("2877"),f=Object(p["a"])(d,o,n,!1,null,"0671cfdd",null);t["default"]=f.exports}}]);