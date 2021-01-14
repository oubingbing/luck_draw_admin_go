var tableHeight = document.body.clientHeight - 120;
new Vue({
    el: '#app',
    data: {
        height:tableHeight,
        total:0,
        page_size:20,
        current_page:1,
        createGift:false,
        giftList:[],
        gift:{
            name:"",
            num:1,
            type:3,
            from:1,
            status:1,
            des:"",
            attachments:[]
        },
        cosToken:"",
        cosDomain:""
    },
    created: function () {
        this.getUserList();
        this.getToken();
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
                    this.$message.error(res.msg);
                }else{
                    let data = [];
                    //this.userList = res.data
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        cancelCreate(done) {
            this.$confirm('确认关闭？')
                .then(_ => {
                    done();
                })
                .catch(_ => {});
        },
        //提交数据
        submitGift(){
            this.createGift = false
        },
        //处理上传前
        uploadPreview(e){

        },
        //处理上传后
        uploadRemove(e){

        },
        getToken:function (e) {
            var url = "/admin/api/cos/token";
            axios.get(url).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.cosToken = res.data.token;
                    this.cosDomain = res.data.domain;
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    }
})