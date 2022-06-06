<template>
    <div class="home">
        <Header></Header>
        <div class="main">
            <LeftMenu></LeftMenu>
            <RightContent></RightContent>
        </div>
    </div>
</template>


<script>
import LeftMenu from '@/components/admin/common/LeftMenu.vue';
import RightContent from '@/components/admin/common/RightContent.vue';
import Header from '@/components/admin/common/Header.vue';

export default {
    name: 'home',
    components: {
        LeftMenu,
        RightContent,
        Header,
    },
    data() {
        return {
        };
    },
    created() {
        // 拿到默认的路由对象,保存全局变量
        this.initTabList();
    },
    computed: {
        // 当前所在模块
        menuModule() {
            return this.$store.state.menuModule;
        },
        // 标签栈
        tabList() {
            return this.$store.state.tabList;
        },
    },
    methods: {
        /**
        *@description: 保存默认的当前模块的首页的路由对象
        *@param{}
        *@return:
        */
        initTabList() {
            // console.log(this.$route.path);
            // 获取匹配到的路由对象
            const firstRoute = this.$router.options.routes.filter(item => item.name === this.menuModule)[0];
            // 循环该路由对象的children
            const { redirect } = firstRoute;
            if (firstRoute && redirect) {
                const saveRoute = firstRoute.children.filter(item => item.path === redirect)[0];
                if (saveRoute) {
                    this.tabList.push(saveRoute);
                    this.$store.commit('tabList', this.tabList);
                }
            }
        },

    },
};
</script>