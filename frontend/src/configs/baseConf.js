const menuList = [
    {
        name: 'GeneralQuery',
        nameZh: "综合查询",
        children: [
            {
                path: '/UseGuide',
                name: 'UseGuide',
                nameZh: "使用指南",
                component: () => import('../views/UseGuide.vue')
            },
            {
                path: '/ResourceApply/UndoList',
                name: 'UndoList',
                nameZh: "待办列表",
                component: () => import('../views/ResourceApply.vue')
            },
            {
                path: '/ResourceApply/DoneList',
                name: 'DoneList',
                nameZh: "已办列表",
                component: () => import('../views/ResourceApply.vue')
            },
        ]
    },
    {
        path: '/ResourceApplication',
        name: 'ResourceApplication',
        nameZh: "资源申请",
        component: () => import('../views/ResourceApplication.vue')
    },
    {
        path: '/ComprehensiveView',
        name: 'ComprehensiveView',
        nameZh: "综合视图",
        component: () => import('../views/ComprehensiveView.vue')
    },
    {
        name: 'MonitoringAlarm',
        nameZh: "监视告警",
        children: [
            {
                path: '/AlarmList',
                name: 'AlarmList',
                nameZh: "告警列表",
                component: () => import('../views/MonitoringAlarm.vue')
            },
            {
                path: '/AlarmSetting',
                name: 'AlarmSetting',
                nameZh: "告警设置",
                component: () => import('../views/AlarmSetting.vue')
            },
            {
                path: '/AlarmStatistical',
                name: 'AlarmStatistical',
                nameZh: "告警统计",
                component: () => import('../views/AlarmStatistical.vue')
            },
        ]
    },
    {
        path: '/StatisticalAnalysis',
        name: 'StatisticalAnalysis',
        nameZh: "统计分析",
        component: () => import('../views/StatisticalAnalysis.vue')
    },
    {
        name: 'SystemManage',
        nameZh: "系统管理",
        children: [
            {
                path: '/UserManage',
                name: 'UserManage',
                nameZh: "用户管理",
                component: () => import('../views/UserManage.vue')
            },
            {
                path: '/SheetManage',
                name: 'SheetManage',
                nameZh: "表单管理",
                component: () => import('../views/SheetManage.vue')
            },
            {
                path: '/AuthorityManagement',
                name: 'AuthorityManagement',
                nameZh: "权限管理",
                component: () => import('../views/AuthorityManagement.vue')
            }
        ]
    }
]

export {
    menuList,
}