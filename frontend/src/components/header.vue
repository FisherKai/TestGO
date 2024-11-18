<template>
  <section class="m-header">
    <section class="AppHeader-globalBar-start">
        <img src="../assets/logo.jpg">
        <p>省级气象基础设施资源池智能运维平台</p>
    </section>
    <section class="AppHeader-globalBar-end"></section>
  </section>
</template>

<script>
export default {

}
</script>

<style scoped>
.m-header{
    height: 6vh;
    display: flex;
    background: linear-gradient(to right, rgb(18,100,196) , rgb(112, 171, 238)); ;
}
img{
    width: 6vh;
    border-radius: 50%;
}
.AppHeader-globalBar-start{
    display: flex;
    justify-content: center;
    align-items: center;
    color: #fff;
}

.AppHeader-globalBar-start p{
    font-size: 22px;
    font-weight: 600;
}
</style>