package com.codex.example.spring.spring;

import org.springframework.data.repository.CrudRepository;
import com.codex.example.spring.spring.User;

public interface UserRepository extends CrudRepository<User,
Integer>{
}
