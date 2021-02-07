var tableHeight = document.body.clientHeight - 120;
var date = formatDate(new Date());
new Vue({
    el: '#app',
    data: {
        operateTitle:"新建活动",
        cosClient:null,
        todayStr:date,
        height:tableHeight,
        total:0,
        page_size:20,
        current_page:1,
        createActivity:false,
        activityList:[],
        gifts:[],
        activity:{
            Id:'',
            name:"",
            type:1,
            gift_id:0,
            limit_join:0,
            join_limit_num:1,
            receive_limit:1,
            des:"",
            attachments:[],
            share_title:"",
            share_image:"",
            big_pic:2,
            draw_type:1,
            really:0,
            open_ad:0,
            is_top:0
        },
        initActivity: {
            name:"",
            type:1,
            gift_id:0,
            limit_join:0,
            join_limit_num:1,
            receive_limit:1,
            des:"",
            attachments:[],
            share_title:"",
            share_image:"",
            big_pic:2,
            draw_type:1,
            really:0,
            open_ad:0,
            is_top:0
        },
        cos:{
            token:"",
            domain:"",
            bucket:"",
            region:"",
            tmp_secret_id:"",
            tmp_secret_key:"",
            start_time:"",
            expired_time:"",
            env:""
        },
        rules: {
            name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
            num: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            type: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            gift_id: [{ required: true, message: '请选择奖品', trigger: 'blur' }],
            status: [{ required: true, message: '请输入数量', trigger: 'blur' }],
            join_limit_num: [{ required: true, message: '请输入限制参与人数', trigger: '' }],
            receive_limit: [{ required: true, message: '请输入限领数量', trigger: '' }],
            attachments: [{ required: true, message: '请上传活动图片', trigger: '' }],
        },
    },
    created: function () {
        this.getPage();
        this.getToken();
        this.getGifts();
    },
    methods: {
        edit:function(row){
            let activity = {
                Id:row.ID,
                name:row.Name,
                type:row.Type,
                gift_id:row.GiftId,
                limit_join:row.LimitJoin,
                join_limit_num:row.JoinLimitNum,
                receive_limit:row.ReceiveLimit,
                des:row.Des,
                attachments:row.Attachments,
                share_title:row.ShareTitle,
                share_image:row.ShareImage,
                big_pic:row.BigPic,
                draw_type:row.DrawType,
                really:row.Really,
                open_ad:row.OpenAd,
                is_top:row.IsTop
            }
            this.activity = activity
            this.createActivity = true
            this.operateTitle = "编辑活动"
        },
        update:function(){
            var url = "/admin/api/activity/update";
            axios.put(url,this.activity).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.$message.success(res.msg);
                    let activityList = this.activityList
                    activityList = activityList.map(item=>{
                        if (item.ID == row.ID){
                            item = res.data
                        }
                        return item
                    })
                    this.activityList = activityList
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        //上架活动
        upActivity:function(id,status){
            var url = "/admin/api/activity/update_status";
            axios.put(url,{id:id,status:status}).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.$message.success(res.msg);
                    let activityList = this.activityList
                    activityList = activityList.map(item=>{
                        if (item.ID == id){
                            item.Status = status
                        }
                        return item
                    })
                    this.activityList = activityList
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        //下架活动
        downActivity:function(id,status){
            var url = "/admin/api/activity/update_status";
            axios.put(url,{id:id,status:status}).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.$message.success(res.msg);
                    let activityList = this.activityList
                    activityList = activityList.map(item=>{
                        if (item.ID == id){
                            item.Status = status
                        }
                        return item
                    })
                    this.activityList = activityList
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        getGifts:function (e) {
            var url = "/admin/api/gift/enable_list";
            axios.get(url,{}).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    let data = [];
                    this.gifts = res.data;
                    if (this.gifts.length > 0){
                        this.activity.gift_id = this.gifts[0]["ID"]
                    }
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        getPage:function (e) {
            var url = "/admin/api/activity/page";
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
                    this.activityList = res.data
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        deleteRow(index, rows) {
            let data = this.activityList[index];
            var url = "/admin/api/activity/delete";
            axios.delete(url,{data:{id:data.ID}}).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    rows.splice(index, 1);
                    this.activityList = rows;
                    this.$message.success(res.msg);
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        cancelCreate(done) {
            this.$confirm('确认关闭？')
                .then(_ => {
                    done();
                })
                .catch(_ => {});
        },

        submit:function(){
            if (this.operateTitle == "新建活动"){
                this.submitActivity()
            }else{
                this.update()
            }
        },

        //提交数据
        submitActivity(){
            var url = "/admin/api/activity/create";
            this.activity.join_limit_num = parseFloat(this.activity.join_limit_num);
            if (this.activity.name == ""){
                this.$message.error("活动名称不能为空");
                return false
            }

            if (this.activity.gift_id <= 0){
                this.$message.error("请选择奖品");
                return false
            }

            if (this.activity.join_limit_num <= 0){
                this.$message.error("限制人数不能为空或者小于等于0");
                return false
            }

            this.activity.receive_limit = parseInt(this.activity.receive_limit)
            if (this.activity.receive_limit <= 0){
                this.$message.error("中奖人数不能为空或者小于等于0");
                return false
            }

            if (this.activity.attachments.length <= 0){
                this.$message.error("活动图片不能为空");
                return false
            }

            if (this.activity.share_image.length <= 0){
                this.activity.share_image = [""]
            }

            axios.post(url,this.activity).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.$message.success(res.msg);
                    this.createActivity = false;
                    this.activity = this.initActivity
                    this.current_page = 1
                    this.activityList = []
                    this.getPage()
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        },
        //上传图片到cos
        upload(e){
            var timestamp=new Date().getTime();
            let fileName = this.cos.env+"/admin/"+this.todayStr+"/"+timestamp+"_"+e.file.name;
            this.cosClient.putObject({
                Bucket: this.cos.bucket,        /* 必须 */
                Region: this.cos.region,        /* 存储桶所在地域，必须字段 */
                Key: fileName,                  /* 必须 */
                StorageClass: 'STANDARD',
                Body: e.file,                   // 上传文件对象
                onProgress: progressData=> {
                    console.log(JSON.stringify(progressData));
                }
            }, (err, data)=> {
                if (err == null){
                    if (e.data == 1){
                        this.activity.share_image = [fileName]
                    }else{
                        this.activity.attachments = [fileName]
                    }
                }else{
                    this.$message.error("图片上传出错");
                }
            });
        },
        //处理上传前
        uploadPreview(e){

        },
        //处理上传后
        uploadRemove(e){
            if (e.data == 1){
                this.activity.share_image = []
            }else{
                this.activity.attachments = []
            }
        },
        getToken:function (e) {
            var url = "/admin/api/cos/token";
            axios.get(url).then( response=> {
                var res = response.data;
                if (res.code != 0){
                    this.$message.error(res.msg);
                }else{
                    this.cos.token          = res.data.token;
                    this.cos.domain         = res.data.domain;
                    this.cos.bucket         = res.data.bucket;
                    this.cos.region         = res.data.region;
                    this.cos.tmp_secret_id  = res.data.tmp_secret_id;
                    this.cos.tmp_secret_key = res.data.tmp_secret_key;
                    this.cos.start_time     = res.data.start_time;
                    this.cos.expired_time   = res.data.expired_ime;
                    this.cos.env            = res.data.env;
                    let cos = this.cos
                    this.cosClient = new COS({
                        // 必选参数
                        getAuthorization: function (options, callback) {
                            callback({
                                TmpSecretId: cos.tmp_secret_id,
                                TmpSecretKey: cos.tmp_secret_key,
                                XCosSecurityToken: cos.token,
                                StartTime: cos.start_time,
                                ExpiredTime: cos.expired_time,
                            });
                        }
                    });
                }
            }).catch(error => {
                this.$message.error("请求异常");
            });
        }
    }
})