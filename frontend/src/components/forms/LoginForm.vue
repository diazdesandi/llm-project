<script setup lang="ts">
import {
  Button, Input, 
  Card, CardContent, CardDescription, CardHeader, CardTitle,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui'
import { toTypedSchema } from '@vee-validate/zod';
import { useForm } from 'vee-validate';
import z from 'zod';


const loginFormSchema = toTypedSchema(z.object({
  email: z
    .string({
      required_error: 'Please enter your email address.',
    })
    .email(),
  password: z
    .string({
      required_error: 'Please enter your password.',
    })
    .min(8, 'Password must be at least 8 characters long.'),
}))


const { handleSubmit, resetForm } = useForm({
  validationSchema: loginFormSchema,
  initialValues: {
    email: 'johndoe@gmail.com',
    password: 'asdfghkhlg',
  },
})

const onSubmit = handleSubmit((values) => {
  console.log('Form submitted:', values)
  // You can replace this with actual submission logic
})


</script>

<template>
  <div class="flex flex-col gap-6">
    <Card>
      <CardHeader class="text-center">
        <CardTitle class="text-xl">
          Welcome back
        </CardTitle>
        <CardDescription>
          Login with your Google account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit="onSubmit">
          <div class="grid gap-6">
            <div class="flex flex-col gap-4">
              <Button variant="outline" class="w-full">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path
                    d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                    fill="currentColor" />
                </svg>
                Login with Google
              </Button>
            </div>
            <div
              class="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
              <span class="relative z-10 bg-background px-2 text-muted-foreground">
                Or continue with
              </span>
            </div>
            <div class="grid gap-6">
              <div class="grid gap-2">
                <FormField v-slot="{ componentField }" name="email">
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input id="email" type="email" placeholder="m@example.com" v-bind:componentField required />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </FormField>
              </div>
              <div class="grid gap-2">
                <div class="">
                  <FormField v-slot="{ componentField }" name="password">
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl class="w-full">
                        <Input id="password" type="password" placeholder="********" v-bind:componentField required />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  </FormField>
                </div>
              </div>
              <Button type="submit" class="w-full" @submit="onSubmit">
                Login
              </Button>
            </div>
            <div class="text-center text-sm">
              Don't have an account?
              <router-link to="/auth/signup" class="underline underline-offset-4">
                Sign up
              </router-link>
            </div>
          </div>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
