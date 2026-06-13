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

const signupFormSchema = toTypedSchema(z.object({
  firstName: z
    .string({
      required_error: 'Please enter your first name.',
    })
    .min(2, 'First name must be at least 2 characters long.'),
  lastName: z
    .string({
      required_error: 'Please enter your last name.',
    })
    .min(2, 'Last name must be at least 2 characters long.'),
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


const { handleSubmit } = useForm({
  validationSchema: signupFormSchema,
  initialValues: {
    firstName: 'Max',
    lastName: 'Robinson',
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
  <Card class="mx-auto max-w-sm">
    <CardHeader>
      <CardTitle class="text-xl">
        Sign Up
      </CardTitle>
      <CardDescription>
        Enter your information to create an account
      </CardDescription>
    </CardHeader>
    <CardContent>
      <form @submit="onSubmit">
        <div class="grid gap-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="grid gap-2">
              <Label for="first-name">First name</Label>
              <Input id="first-name" placeholder="Max" required />
            </div>
            <div class="grid gap-2">
              <Label for="last-name">Last name</Label>
              <Input id="last-name" placeholder="Robinson" required />
            </div>
          </div>
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
          <Button type="submit" class="w-full" @submit="onSubmit">
            Create an account
          </Button>
          <Button variant="outline" class="w-full">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <path
                d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                fill="currentColor" />
            </svg>
            Sign up with Google
          </Button>
        </div>
        <div class="mt-4 text-center text-sm">
          Already have an account?
          <router-link to="/auth/login" class="underline underline-offset-4">
            Sign in
          </router-link>
        </div>
      </form>
    </CardContent>
  </Card>
</template>
