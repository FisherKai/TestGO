<!-- 综合视图 -->
<template>
  <section class="comprehensive-view-container">
    <section class="m-item">
      <span class="m-title" style="font-size: 28px;">华三超融合</span>
      <section class="item">
        <span class="title">资源统计</span>
        <section class="chart-item">
          <p class="sub-title">vCPU:</p>
          <section class="chart-1" ref="chart1"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">内存:</p>
          <section class="chart-1" ref="chart2"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">本地存储:</p>
          <section class="chart-1" ref="chart3"></section>
          <section class="tips"></section>
        </section>
      </section>
      <section class="item">
        <span class="title">实例统计</span>
        <section>
          <section>
            <span>计算实例总数</span>
            <span>68</span>
          </section>
          <section class="chart-2" ref="chart1_1"></section>
        </section>
      </section>
    </section>
    <section class="m-item">
      <span class="m-title" style="font-size: 28px;">华为</span>
      <section class="item">
        <span class="title">资源统计</span>
        <section class="chart-item">
          <p class="sub-title">vCPU:</p>
          <section class="chart-1" ref="chart11"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">内存:</p>
          <section class="chart-1" ref="chart21"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">本地存储:</p>
          <section class="chart-1" ref="chart31"></section>
          <section class="tips"></section>
        </section>
      </section>
      <section class="item">
        <span class="title">实例统计</span>
        <section>
          <section>
            <span>计算实例总数</span>
            <span>143</span>
          </section>
          <section class="chart-2" ref="chart2_1"></section>
        </section>
      </section>
    </section>
    <section class="m-item">
      <p class="m-title" style="font-size: 28px;">VMware</p>
      <section class="item">
        <span class="title">资源统计</span>
        <section class="chart-item">
          <p class="sub-title">vCPU:</p>
          <section class="chart-1" ref="chart12"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">内存:</p>
          <section class="chart-1" ref="chart22"></section>
          <section class="tips"></section>
        </section>
        <section class="chart-item">
          <p class="sub-title">本地存储:</p>
          <section class="chart-1" ref="chart32"></section>
          <section class="tips"></section>
        </section>
      </section>
      <section class="item">
        <span class="title">实例统计</span>
        <section>
          <section>
            <span>计算实例总数</span>
            <span>66</span>
          </section>
          <section class="chart-2" ref="chart3_1"></section>
        </section>
      </section>
    </section>
  </section>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import api from "../api/api.js"
import { onMounted, ref } from "vue";

let chart1 = ref()
let chart2 = ref()
let chart3 = ref()
let chart11 = ref()
let chart21 = ref()
let chart31 = ref()
let chart12 = ref()
let chart22 = ref()
let chart32 = ref()

let chart1_1 = ref()
let chart2_1 = ref()
let chart3_1 = ref()

const getInfo = () =>{
  api.getResourceInfo().then(res => {
    console.log(res)
  });
}

onMounted(() => {
  getInfo()
  const mychart1 = echarts.init(chart1.value)
  const myChart2 = echarts.init(chart2.value)
  const myChart3 = echarts.init(chart3.value)
  const mychart4 = echarts.init(chart11.value)
  const myChart5 = echarts.init(chart21.value)
  const myChart6 = echarts.init(chart31.value)
  const mychart7 = echarts.init(chart12.value)
  const myChart8 = echarts.init(chart22.value)
  const myChart9 = echarts.init(chart32.value)
  
  const mychart10 = echarts.init(chart1_1.value)
  const myChart11 = echarts.init(chart2_1.value)
  const myChart12 = echarts.init(chart3_1.value)

  let opt1 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272', '#fc8251',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1494, name: '已用 1494 个' },
          { value: 960, name: '可用 960 个' },
        ]
      }
    ]
  };
  let opt2 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 3944, name: '已用 3944 GB' },
          { value: 1094, name: '可用 1094 GB' },
        ]
      }
    ]
  };
  let opt3 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: [
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 88, name: '已用 88 TB' },
          { value: 72, name: '可用 72 TB' },
        ]
      }
    ]
  };
  let opt4 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272', '#fc8251',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1182, name: '已用 1182 个' },
          { value: 522, name: '可用 522 个' },
        ]
      }
    ]
  };
  let opt5 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 3639, name: '已用 3639 GB' },
          { value: 1988, name: '可用 1988 GB' },
        ]
      }
    ]
  };
  let opt6 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: [
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 119, name: '已用 119 TB' },
          { value: 17, name: '可用 17 TB' },
        ]
      }
    ]
  };
  let opt7 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272', '#fc8251',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 128, name: '已用 128 个' },
          { value: 512, name: '可用 512 个' },
        ]
      }
    ]
  };
  let opt8 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: ['#3BA272',
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 240, name: '已用 240 GB' },
          { value: 16, name: '可用 16 GB' },
        ]
      }
    ]
  };
  let opt9 = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    color: [
      '#5470c6',
      '#9A60B4',
      '#ef6567',
      '#f9c956',
    ],
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 97, name: '已用 97 TB' },
          { value: 11, name: '可用 11 TB' },
        ]
      }
    ]
  };

  let opt10 =  {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    x:'center',
    y:'top',
    padding: [20,0,0,0]
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
      color: ['#ef6567','#3BA272', '#fc8251',
      '#5470c6',
    ],
  xAxis: {
    type: 'value'
  },
  yAxis: {
    type: 'category',
    data: ['计算实例']
  },
  series: [
    {
      name: '故障',
      type: 'bar',
      barWidth: 50,
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '运行中',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [68]
    },
    {
      name: '已停止',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '其他',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    }
  ]
};

let opt11 =  {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    x:'center',
    y:'top',
    padding: [20,0,0,0]
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
      color: ['#ef6567','#3BA272', '#fc8251',
      '#5470c6',
    ],
  xAxis: {
    type: 'value'
  },
  yAxis: {
    type: 'category',
    data: ['计算实例']
  },
  series: [
    {
      name: '故障',
      type: 'bar',
      barWidth: 50,
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '运行中',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [143]
    },
    {
      name: '已停止',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '其他',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    }
  ]
};

let opt12 =  {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    x:'center',
    y:'top',
    padding: [20,0,0,0]
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
      color: ['#ef6567','#3BA272', '#fc8251',
      '#5470c6',
    ],
  xAxis: {
    type: 'value'
  },
  yAxis: {
    type: 'category',
    data: ['计算实例']
  },
  series: [
    {
      name: '故障',
      type: 'bar',
      barWidth: 50,
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '运行中',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [62]
    },
    {
      name: '已停止',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [0]
    },
    {
      name: '其他',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [4]
    }
  ]
};

  mychart1.setOption(opt1);
  myChart2.setOption(opt2);
  myChart3.setOption(opt3);
  mychart4.setOption(opt4);
  myChart5.setOption(opt5);
  myChart6.setOption(opt6);
  mychart7.setOption(opt7);
  myChart8.setOption(opt8);
  myChart9.setOption(opt9);

  mychart10.setOption(opt10);
  myChart11.setOption(opt11);
  myChart12.setOption(opt12);
})
</script>

<style scoped>
.comprehensive-view-container {
  display: flex;
}

.m-item {
  flex: 1;
}

.chart-item {
  height: 200px;
}

.chart-1 {
  height: 100%;
}

.chart-2{
  height: 180px;
}
.m-title{
  font-size: 22px;
  font-weight: 600;
}
.title {
  font-size: 20px;
  font-weight: 600;
  color: rgba(64, 158, 255, 1);
}

.sub-title {}
</style>