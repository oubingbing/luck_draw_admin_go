
var tableHeight = document.body.clientHeight - 120;
new Vue({
        el: '#app',
        data: {
            height:tableHeight,
            total:0,
            page_size:20,
            current_page:1,
            userList:[]
        },
        created: function () {
            this.getUserList()
        },
        methods: {
            deleteRow(index, rows) {
                rows.splice(index, 1);
            },
            getUserList:function (e) {
                var url = "/admin/api/user";
                axios.get(url+"?page_size="+this.page_size+'&page_num='+this.current_page+'&order_by=created_at&sort=desc',{
                    page_size:this.page_size,
                    page_number:this.current_page,
                    order_by:'created_at',
                    sort_by:'desc'
                }).then( response=> {
                    var res = response.data;
                    if (res.code != 0){

                    }else{
                        let data = [];
                        this.userList = res.data
                        console.log(this.userList)
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            }
        }
    })