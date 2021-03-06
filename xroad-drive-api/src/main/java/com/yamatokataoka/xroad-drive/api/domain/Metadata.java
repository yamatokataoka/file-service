package com.yamatokataoka.xroaddrive.api.domain;

import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.Id;

import java.time.Instant;

public class Metadata {

  @Id
  private String id;
  private String filename;
  private Long filesize;

  @CreatedDate
  private Instant createdDateTime;

  public Metadata() {}

  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }

  public String getFilename() {
    return filename;
  }

  public void setFilename(String filename) {
    this.filename = filename;
  }

  public Long getFilesize() {
    return filesize;
  }

  public void setFilesize(Long filesize) {
    this.filesize = filesize;
  }

  public Instant getCreatedDateTime() {
    return createdDateTime;
  }

  public void setCreatedDateTime(Instant createdDateTime) {
    this.createdDateTime = createdDateTime;
  }
}