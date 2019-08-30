<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12">
        <v-row
          :align="alignment"
          :justify="justify"
          class="grey lighten-5 ml-2"
        >
          <h1 class="display1">About Me</h1>
          <p>
            I am very motivated and work very well with little assistance or guidance due to being self taught in most areas I have worked. I do however work well with others and find great value working on teams and have learned a great deal from current and previous team members. I’m versatile and able to quickly adapt to new situations and environments and work well staying focused under severe pressure. I was a member of the University of Cincinnati’s Track and Field team which taught me not only how to be self motivated and self reliant but to value my teammates and how we pushed each other to do our best making the whole team stronger. I’ve handled IT crisis situations that drove many others around into panic while keeping myself focused and calm until the situation was resolved.
          </p>
          <p>
            I finished working full time as a PhD candidate at the University of Cincinnati majoring in Electrical Engineering. For the first 5 years I also works for U.C. Infrastructure Institute as a research assistant. This work played a role in my efforts to move more into a research oriented field while being able to apply my past work experience into applications of cutting edge technologies that are making their way into the industry.
          </p>
          <p>
            Currently I work as a freelance consultant developing web, mobile and middle-ware apps that bridge the gap between the various open hardware embedded systems I design and build. These systems integrate with existing infrastructure as well as new web and mobile apps to provide various automated solutions to companies that want full control and customization while leveraging the latest cutting edge technologies.  Many implementations involve deploying to multiple cloud providers or VPS/IaaS providers that require integration with their own private cloud infrastructure using a common design pattern and top lever security and privacy where needed.  Hybrid cloud solution integrations are almost always desired which require a more robust and highly customizable and automated infrastructure which we utilize programmable software based infrastructure to streamline the process of on demand reconfigurable development workflows to infrastructure management and expansion as needed and across multiple service providers using the same code base and tools.
          </p>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'about',
  components: {
  }
}
</script>
